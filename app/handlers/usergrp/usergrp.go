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
	// r.Post("/", hr.create)

	// r.Route("/{id}", func(r chi.Router) {
	// 	r.Use(UserCtx)
	// 	r.Delete("/", hr.deleteUser)
	// })

	return r
}

func (h *Handlers) list(w http.ResponseWriter, r *http.Request) {
	// get values from ctx
	// page := web.Param(r, "page")
	// pageNumber, err := strconv.Atoi(page)
	// if err != nil {
	// 	return v1Web.NewRequestError(fmt.Errorf("invalid page format [%s]", page), http.StatusBadRequest)
	// }
	// rows := web.Param(r, "rows")
	// rowsPerPage, err := strconv.Atoi(rows)
	// if err != nil {
	// 	return v1Web.NewRequestError(fmt.Errorf("invalid rows format [%s]", rows), http.StatusBadRequest)
	// }

	page := chi.URLParam(r, "page")
	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		render.Render(w, r, http.StatusBadRequest)
	}
	rows := chi.URLParam(r, "rows")
	rowsPerPage, err := strconv.Atoi(rows)
	if err != nil {
		render.Render(w, r, http.StatusBadRequest)
	}

	users, err := h.User.List(r.Context(), pageNumber, rowsPerPage)
	if err != nil {
		return fmt.Errorf("unable to query for users: %w", err)
	}

	return web.Respond(ctx, w, users, http.StatusOK)
}

// -----------------
//--
// Request and Response payloads for the REST api.
//
// The payloads embed the data model objects an
//
// In a real-world project, it would make sense to put these payloads
// in another file, or another sub-package.
//--

type UserPayload struct {
	*user.User
	Role string `json:"role"`
}

func NewUserPayloadResponse(user *user.User) *UserPayload {
	return &UserPayload{User: user}
}

// Bind on UserPayload will run after the unmarshalling is complete, its
// a good time to focus some post-processing after a decoding.
func (u *UserPayload) Bind(r *http.Request) error {
	return nil
}

func (u *UserPayload) Render(w http.ResponseWriter, r *http.Request) error {
	u.Role = "collaborator"
	return nil
}

// -----------------

// func (h *Handlers) create(w http.ResponseWriter, r *http.Request) {
// 	// get values from ctx
// 	ctx := context.Background()

// 	var nu user.NewUser
// 	if err := web.Decode(r, &nu); err != nil {
// 		fmt.Println(err)
// 	}

// 	usr, err := h.User.Create(nu)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	if err := web.Respond(ctx, w, usr, http.StatusCreated); err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		log.Fatal(err) // TODO ! corrigir
// 		return
// 	}
// }

// func (h *Handlers) deleteUser(w http.ResponseWriter, r *http.Request) {
// 	// TODO claims, auth
// 	// claims, err := auth.GetClaims(ctx)
// 	// if err != nil {>
// 	// 	return v1Web.NewRequestError(auth.ErrForbidden, http.StatusForbidden)
// 	// }

// 	user, ok := r.Context().Value("user").(*user.User)
// 	if !ok {
// 		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
// 	}

// 	// TODO claims, auth
// 	// If you are not an admin and looking to delete someone other than yourself.
// 	// if !claims.Authorized(auth.RoleAdmin) && claims.Subject != userID {
// 	// 	return v1Web.NewRequestError(auth.ErrForbidden, http.StatusForbidden)
// 	// }

// 	if err := h.User.Delete(ctx, userID); err != nil {
// 		switch {
// 		case errors.Is(err, user.ErrInvalidID):
// 			return v1Web.NewRequestError(err, http.StatusBadRequest)
// 		default:
// 			return fmt.Errorf("ID[%s]: %w", userID, err)
// 		}
// 	}

// 	return web.Respond(ctx, w, nil, http.StatusNoContent)
// }

// func UserCtx(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		user := chi.URLParam(r, "id")
// 		ctx := context.WithValue(r.Context(), "user", user)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
