package main

import (
	"embed"

	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	go func() {
		hasUpdate, version, url, err := CheckForUpdate()
		if err == nil && hasUpdate {
			// Покажи уведомление через Wails (или лог)
			fmt.Printf("Доступна новая версия: %s\nСкачать: %s\n", version, url)
			// Можно реализовать показ диалога через frontend
		}
	}()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "gorels",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
