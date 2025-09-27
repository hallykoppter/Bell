package main

import (
	"Bell/api/service"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist all:frontend/out
// all:frontend/out/_next/static/*/* all:frontend/out/_next/static/*/*/*
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := service.NewApp("Bell Sekolah")

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "Bell",
		Width:         900,
		Height:        500,
		Frameless:     true,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
