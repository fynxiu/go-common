package store

import (
	"time"

	"github.com/tidwall/buntdb"
)

var _ Store = &BuntStore{}

type BuntStore struct {
	db *buntdb.DB
}

// NewBuntStore create a Store implemented with BuntDB
func NewBuntStore() (Store, error) {
	db, err := buntdb.Open(":memory:") // Open a file that does not persist to disk.
	if err != nil {
		return nil, err
	}
	return &BuntStore{
		db: db,
	}, nil
}

func (s BuntStore) Set(key string, value string, expireIn time.Duration) error {
	return s.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(
			key,
			value,
			&buntdb.SetOptions{Expires: true, TTL: expireIn},
		)
		return err
	})
}

func (s BuntStore) Get(key string) (string, error) {
	var value string
	err := s.db.View(func(tx *buntdb.Tx) (err error) {
		value, err = tx.Get(key)
		return
	})

	return value, err
}
