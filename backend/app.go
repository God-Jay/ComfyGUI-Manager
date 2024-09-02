package backend

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	goruntime "runtime"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"comfygui-manager/backend/comfyUI"
	"comfygui-manager/backend/output"
	"comfygui-manager/backend/store"
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
	comfyUI.RunProxy(ctx)
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
	cmd := openPath(store.ComfyUIPath)
	return cmd.Start()
}

func openPath(path string) *exec.Cmd {
	var cmd *exec.Cmd
	switch goruntime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", path)
	case "darwin":
		cmd = exec.Command("open", path)
	default:
		cmd = exec.Command("xdg-open", path)
	}
	return cmd
}

func (a *App) OpenFileInDir(dir string, fileName string) error {
	fp := filepath.Join(store.ComfyUIPath, dir, fileName)
	log.Println("open file", fp)
	var cmd *exec.Cmd
	if fileName == "" {
		cmd = openPath(fp)
	} else {
		cmd = selectFile(fp)
	}
	return cmd.Start()
}

func selectFile(fp string) *exec.Cmd {
	var cmd *exec.Cmd
	switch goruntime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", "/select,", fp)
	case "darwin":
		cmd = exec.Command("open", "-R", fp)
	default:
		cmd = exec.Command("xdg-open", fp)
	}
	return cmd
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
