package main

import (
	"embed"
	"github.com/ARUMANDESU/todo-app/internal/domain"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Todo App",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 125, G: 38, B: 4, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		EnumBind: []interface{}{
			domain.AllTaskPriority,
			domain.AllTaskStatus,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
