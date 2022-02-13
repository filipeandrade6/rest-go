package usergrp

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/filipeandrade6/rest-go/internal/core/user"
	"github.com/filipeandrade6/rest-go/pkg/web"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

// https://github.com/go-chi/chi/blob/master/_examples/rest/main.go

type Handlers struct {
	User user.Core
	// Auth *auth.Auth
}

func NewUsrGrp(log *zap.SugaredLogger, db *pgxpool.Pool) http.Handler {
	hr := Handlers{
		User: user.NewCore(log, db),
	}

	r := chi.NewRouter()

	r.Get("/{page}/{rows}", hr.list)

	// r.Route("/{userID}", func(r chi.Router) {
	// 	r.Use(UserCtx) // Load the *User on the request context
	// 	r.Get("/", GetUser)
	// 	r.Put("/", UpdateUser)
	// 	r.Delete("/", DeleteUser)
	// })

	return r
}

func (h *Handlers) list(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		render.Render(w, r, web.ErrInvalidRequest(err))
	}
	rows := chi.URLParam(r, "rows")
	rowsPerPage, err := strconv.Atoi(rows)
	if err != nil {
		render.Render(w, r, web.ErrInvalidRequest(err))
	}

	users, err := h.User.List(r.Context(), pageNumber, rowsPerPage)
	if err != nil {
		render.Render(w, r, web.ErrRender(fmt.Errorf("unable to query for users: %w", err)))
	}

	render.RenderList(w, r, user.NewUsersListResponse(users))
}
