package models

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"

	"comfygui-manager/backend/store"
	"comfygui-manager/backend/util"
)

// File represents a file with its name, path, and size
type File struct {
	Name string `json:"name"`
	Path string `json:"path"`
	// TODO fix marshal unmarshal
	Size    int64     `json:"size"`
	ModTime time.Time `json:"modTime"`
	Sha256  string    `json:"sha256"`
}

func (f File) MarshalJSON() ([]byte, error) {
	type Alias File
	return json.Marshal(&struct {
		Alias
		HumanReadableSize string `json:"size"`
	}{
		Alias:             (Alias)(f),
		HumanReadableSize: util.HumanReadableSize(f.Size),
	})
}

// Dir represents a directory with its name, path, subdirectories, files, number of files, and total size of files
type Dir struct {
	Name          string `json:"name"`
	Path          string `json:"path"`
	parentDir     *Dir
	SubDirs       []*Dir  `json:"subDirs"`
	Files         []*File `json:"files"`
	FileCount     int     `json:"fileCount"`
	TotalFileSize int64   `json:"totalFileSize"`
}

func (d Dir) MarshalJSON() ([]byte, error) {
	type Alias Dir
	return json.Marshal(&struct {
		Alias
		HumanReadableTotalFileSize string `json:"totalFileSize"`
	}{
		Alias:                      (Alias)(d),
		HumanReadableTotalFileSize: util.HumanReadableSize(d.TotalFileSize),
	})
}

func (s *Service) ListModelDir() *Dir {
	modelPath := filepath.Join(store.ComfyUIPath, "models")
	log.Println("listDir", modelPath)
	return listDir(s, modelPath)
}

func listDir(s *Service, modelDir string) *Dir {
	dirMap := make(map[string]*Dir)

	filepath.WalkDir(modelDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			// Create a new Dir struct for the current directory
			newDir := Dir{
				Name: filepath.Base(path),
				Path: path,
			}
			// Add the new directory to its parent directory's subDirs
			parentPath := filepath.Dir(path)
			if parentDir, ok := dirMap[parentPath]; ok {
				newDir.parentDir = parentDir
				parentDir.SubDirs = append(parentDir.SubDirs, &newDir)
			}
			// Add the new directory to the map
			dirMap[path] = &newDir
		} else {
			info, _ := d.Info()
			if info.Size() < 2*util.MB {
				return nil
			}

			var newFile File
			dbFile, _ := s.getDbFile(path)
			if dbFile.ModTime == info.ModTime() {
				newFile = dbFile
				// marshal changed size to string
				newFile.Size = info.Size()
			} else {
				newFile = File{
					Name:    info.Name(),
					Path:    path,
					Size:    info.Size(),
					ModTime: info.ModTime(),
				}
				s.setDbFile(newFile)
			}
			parentPath := filepath.Dir(path)
			if parentDir, ok := dirMap[parentPath]; ok {
				parentDir.Files = append(parentDir.Files, &newFile)
				parentDir.FileCount++
				parentDir.TotalFileSize += newFile.Size
				for pd := parentDir.parentDir; pd != nil; pd = pd.parentDir {
					pd.TotalFileSize += newFile.Size
				}
			}
		}
		return nil
	})

	return dirMap[modelDir]
}

func (s *Service) GetModelFileSHA256(filePath string) string {
	return getFileSHA256(s, filePath)
}

func getFileSHA256(s *Service, filePath string) string {
	dbFile, _ := s.getDbFile(filePath)
	if dbFile.Sha256 != "" {
		return dbFile.Sha256
	}

	log.Println("getFileSHA256 start", filePath)
	start := time.Now()
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	log.Println("getFileSHA256 open file duration:", time.Since(start))
	start = time.Now()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return ""
	}

	log.Println("getFileSHA256 copy duration:", time.Since(start))
	start = time.Now()

	hashString := fmt.Sprintf("%x", hash.Sum(nil))
	log.Println("getFileSHA256 hash duration:", time.Since(start))

	dbFile.Sha256 = hashString
	s.setDbFile(dbFile)

	return hashString
}
