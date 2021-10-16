package main

import (
	"errors"
	"strconv"

	"github.com/boltdb/bolt"
)

type BoltRepository struct {
	db *bolt.DB
}

func NewRepository(db *bolt.DB) *BoltRepository {
	return &BoltRepository{db: db}
}
func (r *BoltRepository) Save(n int, result int) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToBytes(n), intToBytes(result))

	})
}

func (r *BoltRepository) Get(n int) (int, error) {
	var res string

	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data := b.Get(intToBytes(n))
		res = string(data)
		return nil
	})
	if err != nil {
		return "", err
	}

	if res == "" {
		return "", errors.New("token not found")
	}

}

func intToBytes(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
