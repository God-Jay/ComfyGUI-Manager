package models

import (
	"comfygui-manager/backend/store"
	"context"
	"github.com/dgraph-io/badger/v4"
)

type Service struct {
	ctx context.Context
	db  *badger.DB
}

func NewService() *Service {
	return &Service{
		//db: getDb(),
	}
}

func (s *Service) getDbFile(path string) (File, error) {
	file, err := store.GetFromDb[File](path)
	return file, err
}

func (s *Service) setDbFile(file File) error {
	return store.SetDb[File](file.Path, file)
}

func (s *Service) StartUp(ctx context.Context) {
	s.ctx = ctx
}
