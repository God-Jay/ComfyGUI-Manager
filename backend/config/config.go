package config

import (
	"github.com/vrischmann/userdir"
	"path"
)

func GetComfyGUIStorePath() string {
	return path.Join(userdir.GetDataHome(), "ComfyGUI Manager")
}
