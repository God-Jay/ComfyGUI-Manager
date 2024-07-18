package store

import (
	"encoding/json"
	"github.com/dgraph-io/badger/v4"
	"github.com/vrischmann/userdir"
	"log"
	"path"
)

//var db *badger.DB

func getDb() *badger.DB {
	//var err error
	db, err := badger.Open(badger.DefaultOptions(path.Join(userdir.GetDataHome(), "./comfygui.manager.db")).WithValueLogFileSize(1 << 26))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetFromDb[T any](path string) (T, error) {
	db := getDb()
	defer db.Close()
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

func SetDb[T any](path string, t T) error {
	db := getDb()
	defer db.Close()
	return db.Update(func(txn *badger.Txn) error {
		bytes, err := json.Marshal(t)
		if err != nil {
			return err
		}
		e := badger.NewEntry([]byte(path), bytes)
		return txn.SetEntry(e)
	})
}
