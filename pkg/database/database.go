// Package database provides support for access the database.
package database

import (
	"context"
	"errors"
	"net/url"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Set of error variables for CRUD operations.
var (
	ErrDBNotFound        = errors.New("not found")
	ErrDBDuplicatedEntry = errors.New("duplicated entry")
)

// Config is the required properties to use the database.
type Config struct {
	User         string
	Password     string
	Host         string
	Name         string
	MaxIdleConns int
	MaxOpenConns int
	DisableTLS   bool
}

func Connect(cfg Config) (*pgxpool.Pool, error) {
	sslMode := "require"
	if cfg.DisableTLS {
		sslMode = "disable"
	}

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     cfg.Host,
		Path:     cfg.Name,
		RawQuery: q.Encode(),
	}

	pool, err := pgxpool.Connect(context.Background(), u.String())
	if err != nil {
		return nil, err // TODO ErrDBnotfound?
	}

	return pool, nil
}

// TODOs
// StatusCheck returns nil if it can successfully talk to the database. It returns a non-nil error otherwise.
// <struct> Transactor interface needed to begin transaction.
// WithinTran runs passed function and do commit/rollback at the end.
// NamedExecContext is a helper function to execute a CUD operation with logging and tracing.
// NamedQuerySlice is a helper function for executing queries that return a collection of data to be unmarshaled into a slice.
// NamedQueryStruct is a helper function for executing queries that return a single value to be unmarshalled into a struct type.
// queryString provides a pretty print version of the query and parameters.
