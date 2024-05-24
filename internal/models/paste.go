package models

import "time"

type UserPaste struct {
	Text       string    `json:"text"`
	Password   string    `json:"password,omitempty"`
}

type PostgresPaste struct {
	Password  string    `json:"password"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}

type PasswordInput struct {
	Password   string    `json:"password,omitempty"`
}