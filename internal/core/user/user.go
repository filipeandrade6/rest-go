package user

import (
	"errors"
	"fmt"

	"github.com/filipeandrade6/rest-go/internal/core/user/db"
)

type Core struct {
	db db.Store
}

func NewCore(db db.Database) Core {
	return Core{
		db: db.NewStore(db),
	}
}

func (c Core) Create(nu NewUser) (User, error) {
	if err := c.db.Create("1", "abc"); err != nil {
		fmt.Println(err)
		return User{}, errors.New("error")
	}

	return User{Name: "ok"}, nil
}
