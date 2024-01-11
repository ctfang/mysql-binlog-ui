package apps

import (
	"changeme/apps/ctx"
	"context"
	"fmt"
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
	// Perform your teardown here
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
