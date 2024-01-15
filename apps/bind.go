package apps

import (
	"changeme/apps/controllers"
)

func (a *App) GetBind() []interface{} {
	got := controllers.GetAllProvider()

	return got
}
