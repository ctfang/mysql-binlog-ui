package orm

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type UploadLogs struct {
	ID        int64     `gorm:"column:id"`
	File      string    `gorm:"column:file"`
	Status    int64     `gorm:"column:status"`
	Msg       string    `gorm:"column:msg"`
	Timestamp time.Time `gorm:"column:timestamp"`
}

type OrmUploadLogs struct {
	db *gorm.DB
}

func NewOrmUploadLogs() *OrmUploadLogs {
	return &OrmUploadLogs{db: GetDB()}
}

func (receiver *UploadLogs) TableName() string {
	return "upload_logs"
}

func (orm *OrmUploadLogs) TableName() string {
	return "upload_logs"
}

func (orm *OrmUploadLogs) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmUploadLogs) GetTableInfo() interface{} {
	return &UploadLogs{}
}

// Create insert the value into database
func (orm *OrmUploadLogs) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmUploadLogs) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmUploadLogs) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmUploadLogs) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmUploadLogs) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmUploadLogs) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmUploadLogs) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmUploadLogs) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmUploadLogs) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmUploadLogs) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmUploadLogs) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmUploadLogs) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmUploadLogs) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmUploadLogs) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmUploadLogs) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmUploadLogs) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmUploadLogs) Unscoped() *OrmUploadLogs {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmUploadLogs) Insert(row *UploadLogs) error {
	return orm.db.Create(row).Error
}

func (orm *OrmUploadLogs) Inserts(rows []*UploadLogs) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmUploadLogs) Order(value interface{}) *OrmUploadLogs {
	orm.db.Order(value)
	return orm
}

func (orm *OrmUploadLogs) Group(name string) *OrmUploadLogs {
	orm.db.Group(name)
	return orm
}

func (orm *OrmUploadLogs) Limit(limit int) *OrmUploadLogs {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmUploadLogs) Offset(offset int) *OrmUploadLogs {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmUploadLogs) Get() []*UploadLogs {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmUploadLogs) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmUploadLogs) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&UploadLogs{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmUploadLogs) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM common")
}

func (orm *OrmUploadLogs) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmUploadLogs) First(conds ...interface{}) (*UploadLogs, bool) {
	dest := &UploadLogs{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmUploadLogs) Take(conds ...interface{}) (*UploadLogs, int64) {
	dest := &UploadLogs{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmUploadLogs) Last(conds ...interface{}) (*UploadLogs, int64) {
	dest := &UploadLogs{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}
func (orm *OrmUploadLogs) Find(conds ...interface{}) ([]*UploadLogs, int64) {
	list := make([]*UploadLogs, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmUploadLogs) Paginate(page int, limit int) ([]*UploadLogs, int64) {
	var total int64
	list := make([]*UploadLogs, 0)
	orm.db.Count(&total)
	if total > 0 {
		if page == 0 {
			page = 1
		}

		offset := (page - 1) * limit
		tx := orm.db.Offset(offset).Limit(limit).Find(&list)
		if tx.Error != nil {
			logrus.Error(tx.Error)
		}
	}

	return list, total
}

// FindInBatches find records in batches
func (orm *OrmUploadLogs) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmUploadLogs) FirstOrInit(dest *UploadLogs, conds ...interface{}) (*UploadLogs, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmUploadLogs) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmUploadLogs) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmUploadLogs) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmUploadLogs) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmUploadLogs) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmUploadLogs) Where(query interface{}, args ...interface{}) *OrmUploadLogs {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmUploadLogs) Select(query interface{}, args ...interface{}) *OrmUploadLogs {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmUploadLogs) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmUploadLogs) And(fuc func(orm *OrmUploadLogs)) *OrmUploadLogs {
	ormAnd := NewOrmUploadLogs()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmUploadLogs) Or(fuc func(orm *OrmUploadLogs)) *OrmUploadLogs {
	ormOr := NewOrmUploadLogs()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmUploadLogs) WhereId(val int32) *OrmUploadLogs {
	orm.db.Where("`id` = ?", val)
	return orm
}

func (orm *OrmUploadLogs) WhereIdIn(val []int32) *OrmUploadLogs {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmUploadLogs) WhereIdGt(val int32) *OrmUploadLogs {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmUploadLogs) WhereIdGte(val int32) *OrmUploadLogs {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmUploadLogs) WhereIdLt(val int32) *OrmUploadLogs {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmUploadLogs) WhereIdLte(val int32) *OrmUploadLogs {
	orm.db.Where("`id` <= ?", val)
	return orm
}
