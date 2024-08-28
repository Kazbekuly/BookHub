package model

import "time"

type User struct {
	Id           int    `json:"-" db:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordhash"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
