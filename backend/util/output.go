package util

import (
	"comfygui-manager/backend/store"
	"strings"
)

func GetOutputStorePath(outputFileName string) string {
	return strings.Join([]string{store.ComfyUIPath, "output", outputFileName}, ".")
}
