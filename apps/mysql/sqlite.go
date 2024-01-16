package mysql

import (
	"changeme/apps/ctx"
	"changeme/apps/datas"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path"
	"strings"
)

func getSqliteDBName(binlogFilePath string) (dbPath, table string) {
	table = "logs"

	info := path.Base(binlogFilePath)
	dbPath = strings.ReplaceAll(info, "_", "-")
	switch strings.Count(dbPath, ".") {
	case 1:
		arr := strings.Split(info, ".")
		dbPath = arr[0]
		table = arr[1]
	}

	return datas.GetSqlitePath(dbPath), "binlog_" + table
}

// GetDB 链接指定 db 必须手动关闭
func GetDB(dbPath, table string) (*gorm.DB, error) {
	savePath := path.Dir(dbPath)
	if err := os.MkdirAll(savePath, 0755); err != nil {
		return nil, err
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 自定义输出，这里使用标准输出
		logger.Config{
			LogLevel: logger.Warn, // 日志级别，你可以设置为logger.Silent以禁用日志
		},
	)
	db, err := gorm.Open(sqlite.Open(dbPath+".db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		if err != nil {
			return nil, err
		}
	}
	ctx.LogDebug("写入 sqlite = " + dbPath)
	create(db, table)
	return db, err
}

func create(db *gorm.DB, table string) {
	ret := db.Exec(`CREATE TABLE IF NOT EXISTS ` + table + ` (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "schema" TEXT NOT NULL,
  "tables" TEXT NOT NULL,
  "timestamp" TIMESTAMP NOT NULL,
  "event" TEXT NOT NULL,
  "row_1" TEXT NOT NULL,
  "row_2" TEXT
);`)

	if ret.Error != nil {
		fmt.Println(ret.Error.Error())
	}
}
