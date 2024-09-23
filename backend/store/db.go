package store

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/dgraph-io/badger/v4"

	"comfygui-manager/backend/config"
)

var db *badger.DB

func initDB() *badger.DB {
	if db != nil {
		return db
	}
	var err error
	db, err = badger.Open(
		badger.DefaultOptions(config.GetComfyGuiDBPath()).
			WithValueLogFileSize(1 << 28).
			WithBypassLockGuard(true),
	)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetFromDb[T any](path string) (T, error) {
	var t T
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(path))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, &t)
		})
	})
	return t, err
}

func GetFromDbWithPrefix[T any](prefix string) ([]struct {
	Key      string
	FileName string
	Value    T
}, error) {
	var items []struct {
		Key      string
		FileName string
		Value    T
	}
	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = true // 设置为 true 可以预取值
		it := txn.NewIterator(opts)
		defer it.Close()

		// 使用前缀扫描
		for it.Seek([]byte(prefix)); it.ValidForPrefix([]byte(prefix)); it.Next() {
			item := it.Item()
			key := string(item.Key())
			fileName := strings.TrimPrefix(key, prefix+".")
			log.Println("db seek key: ", key, " fileName: ", fileName, " prefix: ", prefix)
			var t T
			err := item.Value(func(val []byte) error {
				return json.Unmarshal(val, &t)
			})
			if err != nil {
				return err
			}
			items = append(items, struct {
				Key      string
				FileName string
				Value    T
			}{
				Key:      key,
				FileName: fileName,
				Value:    t,
			})
		}
		return nil
	})
	return items, err
}

func SetDb[T any](path string, t T) error {
	return db.Update(func(txn *badger.Txn) error {
		bytes, err := json.Marshal(t)
		if err != nil {
			return err
		}
		e := badger.NewEntry([]byte(path), bytes)
		return txn.SetEntry(e)
	})
}

func DelDb(path string) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(path))
	})
}
