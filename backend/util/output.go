package util

import (
	"strings"

	"comfygui-manager/backend/store"
)

type StoreType string

const (
	// default workflow
	StoreTypeWorkflow StoreType = ""
	StoreTypeStar     StoreType = "star"
)

func GetStoreOutputStarPath(outputFileName string) string {
	return getOutputStorePath(outputFileName, StoreTypeStar)
}

func GetStoreOutputWorkflowPath(outputFileName string) string {
	return getOutputStorePath(outputFileName, StoreTypeWorkflow)
}

func getOutputStorePath(outputFileName string, storeType StoreType) string {
	// default, workflow
	if storeType == "" {
		// comfyUIPath.output.filename
		return strings.Join([]string{store.ComfyUIPath, "output", outputFileName}, ".")
	} else {
		// storeType.comfyUIPath.output.filename
		return strings.Join([]string{GetStorePrefix(storeType), outputFileName}, ".")
	}
}

func GetStorePrefix(storeType StoreType) string {
	return strings.Join([]string{string(storeType), store.ComfyUIPath, "output"}, ".")
}
