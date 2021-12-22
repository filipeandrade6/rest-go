package db

import (
	"fmt"

	"github.com/filipeandrade6/rest-go/pkg/database/inmemory"
)

type Store struct {
	db *inmemory.DB
}

func NewStore(db *inmemory.DB) Store {
	return Store{
		db: db,
	}
}

func (s Store) Create(key, value string) error {
	if err := s.db.Create(key, value); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
