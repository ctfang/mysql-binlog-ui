package apps

import (
	"changeme/apps/ctx"
	"changeme/apps/orm"
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// StartUp is called at application startup
func (a *App) StartUp(c context.Context) {
	ctx.SetCtx(c)
	a.ctx = c
}

func (a App) DomReady(ctx context.Context) {
	// Add your action here
}

func (a *App) BeforeClose(ctx context.Context) (prevent bool) {
	return false
}

func (a *App) Shutdown(ctx context.Context) {
	if orm.DB != nil {
		db, err := orm.DB.DB()
		if err == nil {
			db.Close()
		}
	}

	if orm.BinlogDBMap != nil {
		for _, dbTemp := range orm.BinlogDBMap {
			db, err := dbTemp.DB()
			if err == nil {
				db.Close()
			}
		}
	}
}
