package main

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

const (
	dbPath = "resources/my.db"
	bucket = "Redirects"
)

func ConnBoltDB() *bolt.DB {
	db, err := bolt.Open(
		dbPath, 0600,
		&bolt.Options{Timeout: 1 * time.Second},
	)
	if err != nil {
		panic(err)
	}
	initDB(db)
	return db
}

func initDB(db *bolt.DB) {
	paths := make(map[string]string)
	paths["/slices"] = "https://blog.golang.org/go-slices-usage-and-internals"
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(bucket))
		if err != nil {
			fmt.Println("Bucket is already exist")
			return err
		}
		for k, v := range paths {
			err = b.Put([]byte(k), []byte(v))
			if err != nil {
				panic(err)
			}
		}
		return err
	})
}

func GetBoltPaths(db *bolt.DB) map[string]string {
	paths := make(map[string]string)
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return nil
		}
		b.ForEach(func(k []byte, v []byte) error {
			paths[string(k)] = string(v)
			return nil
		})
		return nil
	})
	return paths
}
