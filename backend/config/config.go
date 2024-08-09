package config

import (
	"path"

	"github.com/vrischmann/userdir"
)

func GetComfyGuiDataPath() string {
	return path.Join(userdir.GetDataHome(), "ComfyGUI Manager")
}

func GetComfyGuiDBPath() string {
	return path.Join(GetComfyGuiDataPath(), "comfygui.manager.db")
}
