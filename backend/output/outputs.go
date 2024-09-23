package output

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"comfygui-manager/backend/store"
	"comfygui-manager/backend/util"
)

// TODO change port
var port = 8190

type Output struct {
	ctx        context.Context
	serverCtx  context.Context
	serverStop context.CancelFunc
}

func NewOutput() *Output {
	return &Output{}
}

func (o *Output) StartUp(ctx context.Context) {
	o.ctx = ctx
}

type ImageFile struct {
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
	ModTime time.Time `json:"modTime"`
	Stared  bool      `json:"stared"`
}

func (f ImageFile) MarshalJSON() ([]byte, error) {
	type Alias ImageFile
	return json.Marshal(&struct {
		Alias
		HumanReadableSize string `json:"size"`
	}{
		Alias:             (Alias)(f),
		HumanReadableSize: util.HumanReadableSize(f.Size),
	})
}

// TODO change port
func (o *Output) GetImages() []ImageFile {
	images := make([]ImageFile, 0)

	imageDir := filepath.Join(store.ComfyUIPath, "output")
	filepath.Walk(imageDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" {
				relPath, _ := filepath.Rel(imageDir, path)
				stared, _ := store.GetFromDb[bool](util.GetStoreOutputStarPath(relPath))
				images = append(images, ImageFile{
					Name:    relPath,
					Size:    info.Size(),
					ModTime: info.ModTime(),
					Stared:  stared,
				})
			}
		}
		return nil
	})
	slices.SortFunc(images, func(a, b ImageFile) int {
		return int(b.ModTime.Unix() - a.ModTime.Unix())
	})

	return images
}

// TODO change port
func (o *Output) GetStaredImages() []ImageFile {
	images := make([]ImageFile, 0)

	items, err := store.GetFromDbWithPrefix[bool](util.GetStorePrefix(util.StoreTypeStar))
	if err != nil {
		return nil
	}
	var starMap = make(map[string]bool)
	for _, item := range items {
		starMap[item.FileName] = true
	}

	imageDir := filepath.Join(store.ComfyUIPath, "output")
	filepath.Walk(imageDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" {
				relPath, _ := filepath.Rel(imageDir, path)
				if ok := starMap[relPath]; ok {
					images = append(images, ImageFile{
						Name:    relPath,
						Size:    info.Size(),
						ModTime: info.ModTime(),
						Stared:  true,
					})
				}
			}
		}
		return nil
	})
	slices.SortFunc(images, func(a, b ImageFile) int {
		return int(b.ModTime.Unix() - a.ModTime.Unix())
	})

	return images
}

func (o *Output) StarImg(outputFile string) bool {
	stared, _ := store.GetFromDb[bool](util.GetStoreOutputStarPath(outputFile))
	if stared {
		log.Println("unstar", outputFile)
		store.DelDb(util.GetStoreOutputStarPath(outputFile))
		return false
	} else {
		log.Println("star", outputFile)
		store.SetDb(util.GetStoreOutputStarPath(outputFile), true)
		return true
	}
}

func (o *Output) GetImageWorkflow(outputFile string) string {
	workflow, _ := store.GetFromDb[string](util.GetStoreOutputWorkflowPath(outputFile))
	return workflow
}

func (o *Output) StopImageServer() {
	if o.serverStop != nil {
		o.serverStop()
	}
}

func (o *Output) StartImageServer() {
	if store.ComfyUIPath != "" {
		go o.startImageServer()
	}
}

func (o *Output) startImageServer() {
	o.serverCtx, o.serverStop = context.WithCancel(context.Background())

	mux := http.NewServeMux()
	mux.Handle("/output/", http.StripPrefix("/output/", http.FileServer(http.Dir(filepath.Join(store.ComfyUIPath, "output")))))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	go func() {
		<-o.serverCtx.Done()
		server.Shutdown(context.TODO())
	}()

	err := server.ListenAndServe()
	if err != nil {
		runtime.LogErrorf(o.ctx, "Failed to start server: %v", err)
	}

	o.serverCtx = nil
	o.serverStop = nil
}
