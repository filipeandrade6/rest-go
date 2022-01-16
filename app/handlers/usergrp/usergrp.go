package usergrp

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/filipeandrade6/rest-go/internal/core/user"
	"github.com/filipeandrade6/rest-go/pkg/database/inmemory"
	"github.com/filipeandrade6/rest-go/pkg/web"
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

	r := chi.NewRouter()

	r.Get("/", hr.list)
	r.Post("/", hr.create)
	r.Put("/", hr.update)

	return r
}

func (h *Handlers) list(w http.ResponseWriter, r *http.Request) {
	// get values from ctx
	ctx := context.Background()

	

	if err := h.User.Query(r.Body)

}

func (h *Handlers) create(w http.ResponseWriter, r *http.Request) {
	// get values from ctx
	ctx := context.Background()

	var nu user.NewUser
	if err := web.Decode(r, &nu); err != nil {
		fmt.Println(err)
	}

	usr, err := h.User.Create(nu)
	if err != nil {
		fmt.Println(err)
	}

	if err := web.Respond(ctx, w, usr, http.StatusCreated); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func (h *Handlers) read(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) update(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) delete(w http.ResponseWriter, r *http.Request) {

}
