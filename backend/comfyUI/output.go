package comfyUI

import (
	"comfygui-manager/backend/store"
	"comfygui-manager/backend/util"
)

func (c *ComfyUI) StoreOutput(workflow string, outputFileName string) error {
	return store.SetDb[string](util.GetOutputStorePath(outputFileName), workflow)
}
