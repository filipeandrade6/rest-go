package db

import "time"

type User struct {
	ID           string
	Name         string
	Email        string
	Roles        []string
	PasswordHash []byte
	DateCreated  time.Time
	DateUpdated  time.Time
}

func (u *User) toSlice() []string {
	return [7]string{
		u.ID,
		u.Name,
		u.Email,
		u.Roles[],
		string(u.PasswordHash),
		u.Da
	}
}