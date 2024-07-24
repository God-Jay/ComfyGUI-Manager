package main

import (
	"comfygui-manager/backend"
	"comfygui-manager/backend/comfyUI"
	"comfygui-manager/backend/models"
	"comfygui-manager/backend/output"
	"context"
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"net/http"
	"os"
	"strings"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	comfy := comfyUI.NewComfyUI()
	defer comfy.Shutdown()

	modelSrv := models.NewService()

	op := output.NewOutput()
	defer op.StopImageServer()

	app := backend.NewApp(comfy, op)

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "ComfyGUI Manager",
		Width:     1200,
		Height:    860,
		MinWidth:  1100,
		MinHeight: 800,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: NewFileLoader(),
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.Startup(ctx)
			comfy.StartUp(ctx)
			modelSrv.StartUp(ctx)
			op.StartUp(ctx)
		},
		Bind: []interface{}{
			app,
			comfy,
			modelSrv,
			op,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var err error
	requestedFilename := strings.TrimPrefix(req.URL.Path, "/")
	println("Requesting path:", req.URL.Path)
	println("Requesting file:", requestedFilename)
	fileData, err := os.ReadFile(requestedFilename)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
	}

	res.Write(fileData)
}
