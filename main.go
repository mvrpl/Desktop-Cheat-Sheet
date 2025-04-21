package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:     "Desktop Cheat Sheet",
		Width:     1024,
		Height:    768,
		MinWidth:  1024,
		MinHeight: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewGpuIsDisabled: true,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: false,
                HideTitle:                  false,
                HideTitleBar:               false,
                FullSizeContent:            false,
                UseToolbar:                 false,
                HideToolbarSeparator:       true,
			},
			About: &mac.AboutInfo{
				Title:   "Desktop Cheat Sheet",
				Message: "© mvrpl_br",
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
