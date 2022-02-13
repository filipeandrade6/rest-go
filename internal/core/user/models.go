package user

import (
	"encoding/json"
	"net/http"
	"time"
	"unsafe"

	"github.com/filipeandrade6/rest-go/internal/data/db"
	"github.com/filipeandrade6/rest-go/pkg/web"
	"github.com/go-chi/render"
)

// User represents an individual user.
type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Roles        []string  `json:"roles"`
	PasswordHash []byte    `json:"-"`
	DateCreated  time.Time `json:"date_created"`
	DateUpdated  time.Time `json:"date_updated"`
}

// NewUser contains information needed to create a new User.
type NewUser struct {
	Name            string   `json:"name" validate:"required"`
	Email           string   `json:"email" validate:"required,email"`
	Roles           []string `json:"roles" validate:"required"`
	Password        string   `json:"password" validate:"required"`
	PasswordConfirm string   `json:"password_confirm" validate:"eqfield=Password"`
}

// UpdateUser defines what information may be provided to modify an existing
// User. All fields are optional so clients can send just the fields they want
// changed. It uses pointer fields so we can differentiate between a field that
// was not provided and a field that was provided as explicitly blank. Normally
// we do not want to use pointers to basic types but we make exceptions around
// marshalling/unmarshalling.
type UpdateUser struct {
	Name            *string  `json:"name"`
	Email           *string  `json:"email" validate:"omitempty,email"`
	Roles           []string `json:"roles"`
	Password        *string  `json:"password"`
	PasswordConfirm *string  `json:"password_confirm" validate:"omitempty,eqfield=Password"`
}

// =============================================================================

func toUser(dbUsr db.User) User {
	pu := (*User)(unsafe.Pointer(&dbUsr))
	return *pu
}

func toUserSlice(dbUsrs []db.User) []User {
	users := make([]User, len(dbUsrs))
	for i, dbUsr := range dbUsrs {
		users[i] = toUser(dbUsr)
	}
	return users
}

// =============================================================================

type UserResponse struct {
	*User

	// We add an additional field to the response here.. such as this
	// elapsed computed property
	Elapsed int64 `json:"elapsed"`
}

func NewUserResponse(user *User) *UserResponse {
	return &UserResponse{User: user}
}

func (ur *UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ur); err != nil {
		render.Render(w, r, web.ErrRender(err))
	}
	return nil // TODO pq retorna?
}

func NewUsersListResponse(users []User) []render.Renderer {
	list := []render.Renderer{}
	for _, user := range users {
		list = append(list, NewUserResponse(&user)) // TODO &user? passar por value ou ponteiro?
	}
	return list
}

// type UsersListPayload struct {
// 	Users []*User
// }

// func NewUsersListPayloadResponse(users []*User) *UsersListPayload {
// 	return &UsersListPayload{Users: users}
// }

// func (u *UsersListPayload) Bind(r *http.Request) error {
// 	return nil
// }

// func (u *UsersListPayload) Render(w http.ResponseWriter, r *http.Request) error {
// 	u.Role = "collaborator"
// 	return nil
// }

// =============================================================================
//--
// Request and Response payloads for the REST api.
//
// The payloads embed the data model objects an
//
// In a real-world project, it would make sense to put these payloads
// in another file, or another sub-package.
//--

type UserPayload struct {
	*User
	Role string `json:"role"`
}

func NewUserPayloadResponse(user *User) *UserPayload {
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

// =============================================================================
