package orm

import (
	"changeme/apps/ctx"
	"database/sql"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type Binlog struct {
	ID        int64     `gorm:"column:id"`
	Schema    string    `gorm:"column:schema"`
	Table     string    `gorm:"column:tables"`
	Timestamp time.Time `gorm:"column:timestamp"`
	Event     string    `gorm:"column:event"`
	Row1      string    `gorm:"column:row_1"`
	Row2      string    `gorm:"column:row_2"`

	table string
}

type OrmBinlog struct {
	db *gorm.DB

	database string
	table    string
}

var BinlogDBMap = make(map[string]*gorm.DB)

func NewOrmBinlog(database, table string) *OrmBinlog {
	var db *gorm.DB
	var ok bool
	if db, ok = BinlogDBMap[database]; !ok {
		db = NewDB(database)
		if db == nil {
			return nil
		}
		BinlogDBMap[database] = db
	}

	ctx.LogError(database, table)

	return &OrmBinlog{
		database: database,
		db: db.Model(
			&Binlog{table: table},
		).Table(table),
		table: table,
	}
}

func (receiver *Binlog) TableName() string {
	return receiver.table
}

func (orm *OrmBinlog) TableName() string {
	return orm.table
}

func (orm *OrmBinlog) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmBinlog) GetTableInfo() interface{} {
	return &Binlog{table: orm.table}
}

// Create insert the value into database
func (orm *OrmBinlog) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmBinlog) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmBinlog) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmBinlog) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmBinlog) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmBinlog) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmBinlog) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmBinlog) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmBinlog) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmBinlog) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmBinlog) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmBinlog) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmBinlog) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmBinlog) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmBinlog) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmBinlog) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmBinlog) Unscoped() *OrmBinlog {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmBinlog) Insert(row *Binlog) error {
	return orm.db.Create(row).Error
}

func (orm *OrmBinlog) Inserts(rows []*Binlog) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmBinlog) Order(value interface{}) *OrmBinlog {
	orm.db.Order(value)
	return orm
}

func (orm *OrmBinlog) Group(name string) *OrmBinlog {
	orm.db.Group(name)
	return orm
}

func (orm *OrmBinlog) Limit(limit int) *OrmBinlog {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmBinlog) Offset(offset int) *OrmBinlog {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmBinlog) Get() []*Binlog {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmBinlog) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmBinlog) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&Binlog{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmBinlog) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM common")
}

func (orm *OrmBinlog) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmBinlog) First(conds ...interface{}) (*Binlog, bool) {
	dest := &Binlog{table: orm.table}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmBinlog) Take(conds ...interface{}) (*Binlog, int64) {
	dest := &Binlog{table: orm.table}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmBinlog) Last(conds ...interface{}) (*Binlog, int64) {
	dest := &Binlog{table: orm.table}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}
func (orm *OrmBinlog) Find(conds ...interface{}) ([]*Binlog, int64) {
	list := make([]*Binlog, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmBinlog) Paginate(page int, limit int) ([]*Binlog, int64) {
	var total int64
	list := make([]*Binlog, 0)
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
func (orm *OrmBinlog) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmBinlog) FirstOrInit(dest *Binlog, conds ...interface{}) (*Binlog, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmBinlog) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmBinlog) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmBinlog) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmBinlog) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmBinlog) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmBinlog) Where(query interface{}, args ...interface{}) *OrmBinlog {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmBinlog) Select(query interface{}, args ...interface{}) *OrmBinlog {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmBinlog) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmBinlog) And(fuc func(orm *OrmBinlog)) *OrmBinlog {
	ormAnd := NewOrmBinlog(orm.database, orm.table)
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmBinlog) Or(fuc func(orm *OrmBinlog)) *OrmBinlog {
	ormOr := NewOrmBinlog(orm.database, orm.table)
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmBinlog) WhereId(val int64) *OrmBinlog {
	orm.db.Where("`id` = ?", val)
	return orm
}

func (orm *OrmBinlog) WhereIdIn(val []int64) *OrmBinlog {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmBinlog) WhereIdGt(val int64) *OrmBinlog {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmBinlog) WhereIdGte(val int64) *OrmBinlog {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmBinlog) WhereIdLt(val int64) *OrmBinlog {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmBinlog) WhereIdLte(val int64) *OrmBinlog {
	orm.db.Where("`id` <= ?", val)
	return orm
}
