package db

import "time"

type User struct {
	ID           string    `db:"user_id"`
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	Roles        []string  `db:"roles"`
	PasswordHash []byte    `db:"password_hash"`
	DateCreated  time.Time `db:"date_created"`
	DateUpdated  time.Time `db:"date_updated"`
}
