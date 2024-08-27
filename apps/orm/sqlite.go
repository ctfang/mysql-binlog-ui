package orm

import (
	"changeme/apps/ctx"
	"changeme/apps/datas"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path"
)

var DB *gorm.DB

// GetDB 单例
func GetDB() *gorm.DB {
	if DB != nil {
		return DB
	}
	DB = NewDB(datas.GetPath("/system") + ".db")
	return DB
}

// NewDB 每次都是新的
func NewDB(dbPath string) *gorm.DB {
	savePath := path.Dir(dbPath)
	savePath = datas.ToPath(savePath)
	if err := os.MkdirAll(savePath, 0755); err != nil {
		return nil
	}
	dbPath = datas.ToPath(dbPath)
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: NewSqlLog(logger.Config{
			SlowThreshold: 0,
			LogLevel:      logger.LogLevel(logrus.DebugLevel),
			Colorful:      true,
		}),
	})
	if err != nil {
		ctx.LogError("new DB Open err = ", err)
		return nil
	}

	err = db.AutoMigrate(&UploadLogs{})
	if err != nil {
		ctx.LogError("new DB AutoMigrate err = ", err)
		return nil
	}
	return db
}
