package models

import (
	"time"
)

type User struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type UserStore interface {
	CreateUser(username, password string) (*User, error)
	VerifyPassword(username, password string) (*User, error)
	DeleteUser(username string) error
}
