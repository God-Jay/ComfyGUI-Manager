package backend

import (
	"comfygui-manager/backend/comfyUI"
	"comfygui-manager/backend/output"
	"comfygui-manager/backend/store"
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"os/exec"
	goruntime "runtime"
)

// App struct
type App struct {
	ctx    context.Context
	comfy  *comfyUI.ComfyUI
	output *output.Output
}

// NewApp creates a new App application struct
func NewApp(comfyUI *comfyUI.ComfyUI, op *output.Output) *App {
	return &App{
		comfy:  comfyUI,
		output: op,
	}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	comfyUI.RunProxy()
	a.output.StartImageServer()
}

func (a *App) GetComfyUIPath() string {
	return store.ComfyUIPath
}

func (a *App) setComfyUIPath(comfyUIPath string) error {
	a.comfy.Shutdown()
	a.output.StopImageServer()
	return store.SetComfyUIPath(comfyUIPath)
}

func (a *App) Login(comfyUIPath string) error {
	err := a.setComfyUIPath(comfyUIPath)
	a.output.StartImageServer()
	return err
}

func (a *App) Logout() error {
	return a.setComfyUIPath("")
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

func (a *App) OpenFolder() error {
	var cmd *exec.Cmd
	switch goruntime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", store.ComfyUIPath)
	case "darwin":
		cmd = exec.Command("open", store.ComfyUIPath)
	default:
		cmd = exec.Command("xdg-open", store.ComfyUIPath)
	}
	return cmd.Start()
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
