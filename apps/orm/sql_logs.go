package orm

import (
	"changeme/apps/ctx"
	"gorm.io/gorm/logger"
)

// NewSqlLog skip 需要屏蔽的sql
func NewSqlLog(config logger.Config) logger.Interface {
	got := &ctx.Log{
		Config: config,
	}

	return got
}
