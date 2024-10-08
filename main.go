package main

import (
	"changeme/apps"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"log"
	"runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := apps.NewApp()

	// menu
	appMenu := menu.NewMenu()
	if runtime.GOOS == "darwin" {
		appMenu.Append(menu.AppMenu())
		appMenu.Append(menu.EditMenu())
		appMenu.Append(menu.WindowMenu())
	}

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "mysql-binlog",
		Width:             1280,
		Height:            768,
		MinWidth:          1024,
		MinHeight:         768,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         runtime.GOOS != "darwin",
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  options.NewRGBA(27, 38, 54, 0),
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Menu:                     appMenu,
		EnableDefaultContextMenu: true,
		Logger:                   nil,
		LogLevel:                 logger.DEBUG,
		OnStartup:                app.StartUp,
		OnDomReady:               app.DomReady,
		OnBeforeClose:            app.BeforeClose,
		OnShutdown:               app.Shutdown,
		WindowStartState:         options.Normal,
		Bind:                     app.GetBind(),
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			DisableFramelessWindowDecorations: true,
		},
		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHiddenInset(),
			WebviewIsTransparent: false,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "mysql-binlog-ui",
				Message: "一个mysql binlog解析工具",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
