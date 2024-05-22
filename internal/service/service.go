package service

import (
	"pastebin/internal/models"
	"pastebin/internal/repository"
)

type Service struct {
	Paste
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Paste: NewPasteService(repos),
	}
}

type Paste interface {
	Create(userPaste models.UserPaste) (string, error)
	Get(id string, password string) (string, error)
}