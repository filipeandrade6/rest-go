package db

import (
	"context"
	"fmt"

	"github.com/ardanlabs/service/business/sys/database"
)

type Database interface {
	Create(key, value string) error
	Read(key string) (string, error)
	Update(key, value string) error
	Delete(key string) error
}

type Store struct {
	// log
	// transactor
	// isWithinTran
	db Database
}

func NewStore(db Database) Store {
	return Store{
		db: db,
	}
}

// Create inserts a new user into the database.
func (s Store) Create(ctx context.Context, usr User) error {
	const q = `
	INSERT INTO users
		(user_id, name, email, password_hash, roles, date_created, date_updated)
	VALUES
		(:user_id, :name, :email, :password_hash, :roles, :date_created, :date_updated)
		($1, $2, $3, $4, $5, $6, $7)`

	params = [7]string{usr.ID, usr.Name, usr.Email, string(usr.PasswordHash) /* ,*/, usr.DateCreated, usr.DateUpdated}

	if err := database.ExecContext(ctx /*s.log,*/, s.db, q, usr); err != nil {
		return fmt.Errorf("inserting user: %w", err)
	}

	return nil
}

// Update replaces a user document in the database.
func (s Store) Update(ctx context.Context, usr User) error {
	const q = `
	UPDATE
		users
	SET
		"name" = :name,
		"email" = :email,
		"roles" = :roles,
		"password_hash" = :password_hash,
		"date_updated" = :date_updated
	WHERE
		user_id = :user_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, usr); err != nil {
		return fmt.Errorf("updating userID[%s]: %w", usr.ID, err)
	}

	return nil
}

// Delete removes a user from the database.
func (s Store) Delete(ctx context.Context, userID string) error {
	data := struct {
		UserID string `db:"user_id"`
	}{
		UserID: userID,
	}

	const q = `
	DELETE FROM
		users
	WHERE
		user_id = :user_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("deleting userID[%s]: %w", userID, err)
	}

	return nil
}

// Query retrieves a list of existing users from the database.
func (s Store) Query(ctx context.Context, pageNumber int, rowsPerPage int) ([]User, error) {
	data := struct {
		Offset      int `db:"offset"`
		RowsPerPage int `db:"rows_per_page"`
	}{
		Offset:      (pageNumber - 1) * rowsPerPage,
		RowsPerPage: rowsPerPage,
	}

	const q = `
	SELECT
		*
	FROM
		users
	ORDER BY
		user_id
	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY`

	var usrs []User
	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &usrs); err != nil {
		return nil, fmt.Errorf("selecting users: %w", err)
	}

	return usrs, nil
}

// QueryByID gets the specified user from the database.
func (s Store) QueryByID(ctx context.Context, userID string) (User, error) {
	data := struct {
		UserID string `db:"user_id"`
	}{
		UserID: userID,
	}

	const q = `
	SELECT
		*
	FROM
		users
	WHERE
		user_id = :user_id`

	var usr User
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &usr); err != nil {
		return User{}, fmt.Errorf("selecting userID[%q]: %w", userID, err)
	}

	return usr, nil
}

// QueryByEmail gets the specified user from the database by email.
func (s Store) QueryByEmail(ctx context.Context, email string) (User, error) {
	data := struct {
		Email string `db:"email"`
	}{
		Email: email,
	}

	const q = `
	SELECT
		*
	FROM
		users
	WHERE
		email = :email`

	var usr User
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &usr); err != nil {
		return User{}, fmt.Errorf("selecting email[%q]: %w", email, err)
	}

	return usr, nil
}
