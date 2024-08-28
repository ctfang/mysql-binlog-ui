package controllers

import (
	"changeme/apps/ctx"
	"changeme/apps/orm"
	"errors"
	"time"
)

// Binlog @Bean
type Binlog struct{}

func (b Binlog) GetTitleList(id int64) ([]*orm.UploadLogs, error) {
	data, ok := orm.NewOrmUploadLogs().WhereId(id).First()
	if !ok {
		return nil, errors.New("data not is exist")
	}

	dbList := orm.NewOrmUploadLogs().WhereDatabase(data.Database).Limit(100).Order("id DESC").Get()
	return dbList, nil
}

type SearchBinlog struct {
	Table string
	Event string
	Text  string
	Date  []string

	Page  int
	Limit int
}

type DetailsList struct {
	List  []*orm.Binlog
	Total int64
}

func (b Binlog) GetDetailsList(id int64, search *SearchBinlog) (*DetailsList, error) {
	data, ok := orm.NewOrmUploadLogs().WhereId(id).First()
	if !ok {
		return nil, errors.New("data not is exist")
	}

	database, table := data.Database, data.Table
	model := orm.NewOrmBinlog(database, table).Order("id DESC").Order("id DESC")
	if search.Table != "" {
		model.Where("tables = ?", search.Table)
	}
	if search.Event != "" {
		model.Where("event = ?", search.Event)
	}
	if search.Text != "" {
		model.Where("(row_1 like ? or row_2 like ?)", "%"+search.Text+"%", "%"+search.Text+"%")
	}

	if search.Date != nil && len(search.Date) == 2 {
		if search.Date[0] != "" && search.Date[1] != "" {
			start, err := ParseISODateTime(search.Date[0])
			if err != nil {
				ctx.LogError("start date err := ", err)
			}
			end, err := ParseISODateTime(search.Date[1])
			if err != nil {
				ctx.LogError("end date err := ", err)
			}

			model.Where("(timestamp >= ? and timestamp <= ?)", start, end)
		}
	}

	list, count := model.Paginate(search.Page, search.Limit)

	return &DetailsList{
		List:  list,
		Total: count,
	}, nil
}

// ParseISODateTime 将 ISO 8601 格式的时间字符串转换为 time.Time 对象
func ParseISODateTime(dateTimeStr string) (time.Time, error) {
	layout := "2006-01-02T15:04:05.000Z"
	return time.Parse(layout, dateTimeStr)
}
