package db

import (
	"github.com/dgraph-io/badger/v3"
	"sync"
)

var db *badger.DB
var dbOnce = sync.Once{}

func initDb() {
	open, err := badger.Open(badger.DefaultOptions("persist"))
	if err != nil {
		panic(err)
	}
	db = open
}

func GetDb() *badger.DB {
	dbOnce.Do(initDb)
	return db
}

func Set(db *badger.DB, key []byte, value []byte) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
}

func Get(db *badger.DB, key []byte) ([]byte, error) {
	var value []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		value, err = item.ValueCopy(nil)
		return err
	})
	return value, err
}

func AtomicGetSet(db *badger.DB, key []byte, value []byte) ([]byte, error) {
	var oldValue []byte
	err := db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		oldValue, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}
		return txn.Set(key, value)
	})
	return oldValue, err
}

func Keys(db *badger.DB, prefix []byte) ([][]byte, error) {
	var keys [][]byte
	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			key := item.KeyCopy(nil)
			keys = append(keys, key)
		}
		return nil
	})
	return keys, err
}

func Exist(db *badger.DB, key []byte) (bool, error) {
	var exist bool
	err := db.View(func(txn *badger.Txn) error {
		_, err := txn.Get(key)
		if err != nil {
			return err
		}
		exist = true
		return nil
	})
	return exist, err
}
