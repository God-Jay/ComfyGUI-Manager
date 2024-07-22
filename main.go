package main

import (
	"comfygui-manager/backend"
	"comfygui-manager/backend/comfyUI"
	"comfygui-manager/backend/models"
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	comfy := comfyUI.NewComfyUI()
	defer comfy.Shutdown()

	app := backend.NewApp(comfy)

	modelSrv := models.NewService()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "comfygui-manager",
		Width:     1200,
		Height:    860,
		MinWidth:  1100,
		MinHeight: 800,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.Startup(ctx)
			comfy.StartUp(ctx)
			modelSrv.StartUp(ctx)
		},
		Bind: []interface{}{
			app,
			comfy,
			modelSrv,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
