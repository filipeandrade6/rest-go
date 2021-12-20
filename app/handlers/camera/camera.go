package camera

import (
	"errors"
	"fmt"

	"github.com/filipeandrade6/rest-go/app/infile"
	"go.uber.org/zap"
)

type Camera struct {
	log *zap.SugaredLogger
	db  *infile.DB
}

func NewCamera(log *zap.SugaredLogger, db *infile.DB) *Camera {
	return &Camera{
		log: log,
		db:  db,
	}
}

func (c *Camera) QueryByID(input string) (string, error) {

	fmt.Println(input)

	r, err := c.db.Read(input)
	if err != nil {
		return "", errors.New("doesn't exist")
	}

	return r, nil
}
