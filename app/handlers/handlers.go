package handlers

import (
	"net/http"
	"time"

	"github.com/filipeandrade6/rest-go/app/handlers/usergrp"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v4"
	"github.com/treastech/logger"
	"go.uber.org/zap"
)

func NewAPI(log *zap.SugaredLogger, db *pgx.Conn) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(logger.Logger(log.Desugar()))         // ! utilizar?
	r.Use(middleware.Timeout(60 * time.Second)) // * pensar em um timeout.

	// Protected routes
	r.Group(func(r chi.Router) {
		// r.Use(jwtauth.Verifier(tokenAuth))
		// NOTE jwtauth.Authenticator should be added
		// by different router. For example, some routes
		// allow GET and disallow POST

		r.Mount("/users", usergrp.NewUsrGrp(log, db)) // * passar o tokenAuth como dependencia.
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("hello world!"))
		})
	})

	return r
}
