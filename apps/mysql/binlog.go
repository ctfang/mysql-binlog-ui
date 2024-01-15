package mysql

import (
	"changeme/apps/ctx"
	"changeme/apps/datas"
	"changeme/apps/orm"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-mysql-org/go-mysql/replication"
	"gorm.io/gorm"
	"os"
	"sync"
	"time"
)

var logsChan = make(chan *BinlogData, 1000)
var logsChanStatus = make(chan bool)

// BinlogData represents the structure of the "binlog_data" table.
type BinlogData struct {
	ID        int64     `gorm:"column:id"`
	Schema    string    `gorm:"column:schema"`
	Table     string    `gorm:"column:tables"`
	Timestamp time.Time `gorm:"column:timestamp"`
	Event     string    `gorm:"column:event"`
	Row1      string    `gorm:"column:row_1"`
	Row2      string    `gorm:"column:row_2"`
}

var lock sync.Mutex

var DecodeRowCount uint64 = 0

func SaveToSqlite(binlogFilePath string) error {
	if !lock.TryLock() {
		return errors.New("有数据在解析中")
	}
	defer lock.Unlock()

	DecodeRowCount = 0
	go watchLogsChan(binlogFilePath)
	defer func() {
		logsChanStatus <- true
	}()

	var data *BinlogData

	f := func(e *replication.BinlogEvent) error {
		// 处理不同类型的事件
		switch ee := e.Event.(type) {
		case *replication.RotateEvent:
			fmt.Println("RotateEvent:", ee.NextLogName, ee.Position)
		case *replication.RowsEvent:
			data = &BinlogData{
				Schema:    string(ee.Table.Schema),
				Table:     string(ee.Table.Table),
				Timestamp: time.Unix(int64(e.Header.Timestamp), 0),
			}
			switch e.Header.EventType {
			case replication.WRITE_ROWS_EVENTv1, replication.WRITE_ROWS_EVENTv2:
				data.Event = "insert"
				if len(ee.Rows) == 1 {
					data.Event = "insert"
				} else {
					data.Event = "inserts"
				}
				r1, _ := json.Marshal(ee.Rows)
				data.Row1 = string(r1)
			case replication.UPDATE_ROWS_EVENTv1, replication.UPDATE_ROWS_EVENTv2:
				data.Event = "update"
				if len(ee.Rows) == 2 {
					r1, _ := json.Marshal(ee.Rows[0])
					r2, _ := json.Marshal(ee.Rows[1])
					data.Row1 = string(r1)
					data.Row2 = string(r2)
				}
			case replication.DELETE_ROWS_EVENTv1, replication.DELETE_ROWS_EVENTv2:
				if len(ee.Rows) == 1 {
					data.Event = "delete"
				} else {
					data.Event = "deletes"
				}
				r1, _ := json.Marshal(ee.Rows)
				data.Row1 = string(r1)
			}

			DecodeRowCount++
			// 插入数据
			logsChan <- data
		case *replication.QueryEvent:
		case *replication.FormatDescriptionEvent: // 用来描述这些数字代表的事件类型的具体格式和定义
		case *replication.PreviousGTIDsEvent: // 记录上一个GTID的信息。GTID是一个全局唯一的事务标识符，用于在复制环境中准确地追踪数据的复制进度
		case *replication.GTIDEvent: // GTIDEvent会在每个事务开始时记录，并包含当前事务的GTID信息，以便在复制过程中进行数据同步
		case *replication.TableMapEvent: // TableMapEvent用于提供每个表的结构信息
		case *replication.XIDEvent: // XIDEvent用于标识这个事务的提交点
			//ee.Dump(os.Stdout)
		default:
			ee.Dump(os.Stdout)
		}
		return nil
	}
	// 创建一个BinlogReader来读取binlog事件
	r := replication.NewBinlogParser()
	err := r.ParseFile(binlogFilePath, 0, f)
	if err != nil {
		ctx.LogError("ParseFile err = " + err.Error())
		return err
	}

	return nil
}

func watchLogsChan(binlogFilePath string) {
	dbPath, table := getSqliteDBName(binlogFilePath)
	db, err := GetDB(dbPath, table)
	if err != nil {
		ctx.LogError("watchLogsChan db err = " + err.Error() + " path = " + binlogFilePath)
		return
	}
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	logData := &orm.UploadLogs{
		Database:  dbPath + ".db",
		Table:     table,
		File:      binlogFilePath,
		FileSize:  datas.GetFileSizeMB(binlogFilePath),
		Timestamp: time.Now(),
	}
	err = orm.NewOrmUploadLogs().Insert(logData)
	if err != nil {
		ctx.LogError("写入 logs 失败, err = ", err)
		return
	}

	insertDatas := make([]*BinlogData, 0)
	insertCount := 0

	loop := true
	for loop {
		select {
		case data := <-logsChan:
			insertDatas = append(insertDatas, data)
			insertCount++
			if insertCount >= 50 {
				insertCount = 0
				addToSqlite(db, table, insertDatas)
				insertDatas = make([]*BinlogData, 0)
			}
		case <-logsChanStatus:
			insertCount = 0
			addToSqlite(db, table, insertDatas)
			ctx.LogDebug("完成 " + binlogFilePath)
			loop = false
			break
		}
	}

	logData.Status = 1
	logData.Msg = "运行完成"
	orm.NewOrmUploadLogs().WhereId(logData.ID).Updates(logData)
}

func addToSqlite(db *gorm.DB, table string, data []*BinlogData) {
	ret := db.Table(table).CreateInBatches(data, 100)
	if ret.Error != nil {
		ctx.LogError("addToSqlite err = " + ret.Error.Error())
	}
}
