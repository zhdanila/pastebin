package models

type UserPaste struct {
	Text       string    `json:"text"`
	Password   string    `json:"password,omitempty"`
}