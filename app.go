package main

import (
	"comfygui-manager/comfyUI"
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

// App struct
type App struct {
	ctx         context.Context
	ComfyUIPath string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	comfyUI.RunProxy()
}

func (a *App) SetComfyUIPath() {
}

func (a *App) SelectFolder(title string) string {
	selection, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: title,
	})
	if err != nil {
		return ""
	}
	return selection
}

func (a *App) Prompt(defaultFileName, title string) string {
	dialog, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultDirectory:           "",
		DefaultFilename:            defaultFileName,
		Title:                      title,
		Filters:                    nil,
		ShowHiddenFiles:            false,
		CanCreateDirectories:       false,
		TreatPackagesAsDirectories: false,
	})
	if err != nil {
		return ""
	}
	return dialog
}

func (a *App) Confirm(title, msg string) string {
	dialog, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          "",
		Title:         title,
		Message:       msg,
		Buttons:       nil,
		DefaultButton: "",
		CancelButton:  "",
		Icon:          nil,
	})
	if err != nil {
		return ""
	}
	return dialog
}

func (a *App) SaveFile(filePath string, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}
