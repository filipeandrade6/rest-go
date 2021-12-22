package inmemory

import "errors"

type DB map[string]string

func New() *DB {
	db := make(DB)
	return &db
}

func (db DB) Create(key, value string) error {
	if _, ok := db[key]; ok {
		return errors.New(key + " already exists")
	}

	db[key] = value

	return nil
}

func (db DB) Read(key string) (string, error) {
	value, ok := db[key]
	if !ok {
		return "", errors.New(key + " doesn't exists")
	}

	return value, nil
}

func (db DB) Update(key, value string) error {
	if _, ok := db[key]; !ok {
		return errors.New(key + " doesn't exists")
	}

	db[key] = value

	return nil
}

func (db DB) Delete(key string) error {
	if _, ok := db[key]; !ok {
		return errors.New(key + " doesn't exists")
	}

	delete(db, key)

	return nil
}
