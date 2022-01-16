package db

import (
	"fmt"
)

type Database interface {
	Create(key, value string) error
	Read(key string) (string, error)
	Update(key, value string) error
	Delete(key string) error
}

type Store struct {
	db Database
}

func NewStore(db Database) Store {
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
