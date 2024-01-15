package ctx

import (
	"context"
	"errors"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

func LogDebug(msg string, arr ...interface{}) {
	if ctx == nil {
		fmt.Println(msg + fmt.Sprint(arr...))
		return
	}
	if len(arr) > 0 {
		msg = msg + fmt.Sprint(arr...)
	}
	runtime.LogDebug(ctx, msg)
}

func LogInfo(msg string, arr ...interface{}) {
	if ctx == nil {
		fmt.Println(msg + fmt.Sprint(arr...))
		return
	}
	if len(arr) > 0 {
		msg = msg + fmt.Sprint(arr...)
	}
	runtime.LogInfo(ctx, msg)
}

func LogWarn(msg string, arr ...interface{}) {
	if ctx == nil {
		fmt.Println(msg + fmt.Sprint(arr...))
		return
	}
	if len(arr) > 0 {
		msg = msg + fmt.Sprint(arr...)
	}
	runtime.LogWarning(ctx, msg)
}

func LogError(msg string, arr ...interface{}) {
	if ctx == nil {
		fmt.Println(msg + fmt.Sprint(arr...))
		return
	}
	if len(arr) > 0 {
		msg = msg + fmt.Sprint(arr...)
	}
	runtime.LogError(ctx, msg)
}

type Log struct {
	logger.Config
}

func (l *Log) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *Log) Info(ctx2 context.Context, s string, i ...interface{}) {
	if l.LogLevel >= logger.Info {
		LogInfo(s, append([]interface{}{utils.FileWithLineNum()}, i...))
	}
}

func (l *Log) Warn(ctx2 context.Context, s string, i ...interface{}) {
	if l.LogLevel >= logger.Warn {
		LogWarn(s, append([]interface{}{utils.FileWithLineNum()}, i...))
	}
}

func (l *Log) Error(ctx2 context.Context, s string, i ...interface{}) {
	if l.LogLevel >= logger.Error {
		LogError(s, append([]interface{}{utils.FileWithLineNum()}, i...))
	}
}

func (l *Log) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}
	switch {
	case err != nil && !errors.Is(err, logger.ErrRecordNotFound):
		sql, _ := fc()
		LogDebug(fmt.Sprintf("%s\n%s", err.Error(), sql))
	case l.LogLevel >= logger.Info:
		sql, _ := fc()
		LogDebug(sql)
	}
}
