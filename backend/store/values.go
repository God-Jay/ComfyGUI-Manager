package store

var _ = initDB()

var ComfyUIPath = getComfyUIPath()

func getComfyUIPath() string {
	comfyUIPath, _ := GetFromDb[string](KeyComfyUIPath)
	return comfyUIPath
}

func SetComfyUIPath(comfyUIPath string) error {
	err := SetDb[string](KeyComfyUIPath, comfyUIPath)
	if err == nil {
		ComfyUIPath = comfyUIPath
	}
	return err
}
