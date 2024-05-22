package models

import "time"

type UserPaste struct {
	Text       string    `json:"text"`
	Password   string    `json:"password,omitempty"`
}

type PostgresPaste struct {
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}

type PasswordInput struct {
	Password   string    `json:"password,omitempty"`
}