package orm

import (
	"changeme/apps/datas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	if DB != nil {
		return DB
	}
	dbPath := datas.GetPath("/system")

	savePath := path.Dir(dbPath)
	if err := os.MkdirAll(savePath, 0755); err != nil {
		return nil
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
			return nil
		}
	}

	err = db.AutoMigrate(&UploadLogs{})
	if err != nil {
		return nil
	}
	return db
}
