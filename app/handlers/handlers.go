package handlers

import (
	"net/http"
	"time"

	"github.com/filipeandrade6/rest-go/app/handlers/usergrp"
	"github.com/filipeandrade6/rest-go/pkg/database/inmemory"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/treastech/logger"
	"go.uber.org/zap"
)

func NewAPI(log *zap.SugaredLogger, db *inmemory.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(logger.Logger(log.Desugar()))                  // ! utilizar?
	r.Use(render.SetContentType(render.ContentTypeJSON)) // ! utilizar?
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", list)

	// Protected routes
	r.Mount("/users", usergrp.NewUsrGrp(db))
	r.Mount("/abc", routerer(db))

	// TODO userStore := userstore.New(log?, db?)
	// TODO r.Mount("/users", usersresource.New(userStore, tokenAuth <- passado como dependencia de NewRouter))

	// Public routes
	r.Group(func(r chi.Router) {})

	return r
}

func list(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi!"))
}

func routerer(db *inmemory.DB) http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi!"))
	})

	return r
}
