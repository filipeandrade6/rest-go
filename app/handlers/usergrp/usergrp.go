package usergrp

import (
	"net/http"

	"github.com/filipeandrade6/rest-go/internal/core/user"
	"github.com/filipeandrade6/rest-go/pkg/database/inmemory"
	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	User user.Core
	// Auth *auth.Auth
}

func NewUsrGrp(db *inmemory.DB) http.Handler {
	hr := Handlers{
		User: user.NewCore(db),
	}

	return hr.route()
}

func (h *Handlers) route() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.list)

	return r
}

func (h *Handlers) list(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi!"))
}

// func (h Handlers) Create(w http.ResponseWriter, r *http.Request) error { // ! return error?
// 	// get "web" values from context

// 	nu := user.NewUser{}
// 	_, err := h.User.Create(nu)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	w.Write([]byte("hi"))

// 	return nil
// }
