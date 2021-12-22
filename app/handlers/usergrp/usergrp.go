package usergrp

import (
	"fmt"
	"net/http"

	"github.com/filipeandrade6/rest-go/internal/core/user"
)

type Handlers struct {
	User user.Core
	// Auth *auth.Auth
}

func (h Handlers) Create(w http.ResponseWriter, r *http.Request) error {
	// get "web" values from context

	nu := user.NewUser{}
	_, err := h.User.Create(nu)
	if err != nil {
		fmt.Println(err)
		return err
	}

	w.Write([]byte("hi"))

	return nil
}
