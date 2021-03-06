// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"

	"github.com/google/uuid"
)

type Product struct {
	ProductID   uuid.UUID
	Name        sql.NullString
	Cost        sql.NullInt32
	Quantity    sql.NullInt32
	UserID      uuid.NullUUID
	DateCreated sql.NullTime
	DateUpdated sql.NullTime
}

type Sale struct {
	SaleID      uuid.UUID
	UserID      uuid.NullUUID
	ProductID   uuid.NullUUID
	Quantity    sql.NullInt32
	Paid        sql.NullInt32
	DateCreated sql.NullTime
}

type User struct {
	UserID       uuid.UUID
	Name         sql.NullString
	Email        sql.NullString
	Roles        []string
	PasswordHash sql.NullString
	DateCreated  sql.NullTime
	DateUpdated  sql.NullTime
}
