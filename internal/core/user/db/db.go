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

func (s Store) Read(key string) (string, error) {
	id, err := s.db.Read(key) // TODO arrumar
	if err != nil {
		return "", err
	}

	usr := &User{ID: id}

	return usr, nil
}

func (s Store) Delete(key string) error {
	if err := s.db.Delete(key); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
