package user

import (
	"errors"
	"fmt"

	"github.com/filipeandrade6/rest-go/internal/core/user/db"
	"github.com/filipeandrade6/rest-go/pkg/database/inmemory"
)

type Core struct {
	// db db.Store
	db db.Database
}

func NewCore(db db.Database) Core {
	return Core{
		// db: db.NewStore(db),
		db: inmemory.New(),
	}
}

func (c Core) Create(nu NewUser) (User, error) {
	if err := c.db.Create("1", "abc"); err != nil {
		fmt.Println(err)
		return User{}, errors.New("error")
	}

	return User{Name: "ok"}, nil
}
