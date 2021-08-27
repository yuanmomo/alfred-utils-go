package common

import (
	"github.com/boltdb/bolt"
	"log"
	"time"
)

type BoltDB struct {
	db *bolt.DB
}

func NewBoltDB(path string) *BoltDB {
	db, err := bolt.Open(path, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return &BoltDB{db}
}

func (b *BoltDB) Path() string {
	return b.db.Path()
}

func (b *BoltDB) Close() {
	b.db.Close()
}