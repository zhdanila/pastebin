package service

import (
	"pastebin/internal/models"
	"pastebin/internal/repository"
)

type PasteService struct {
	repo *repository.Repository
}

func NewPasteService(repo *repository.Repository) *PasteService {
	return &PasteService{repo: repo}
}

func(s *PasteService) Create(userPaste models.UserPaste) (string, error) {
	return s.repo.Paste.Create(userPaste)
}

func(s *PasteService) Get(id string, password string) (string, error) {
	return s.repo.Paste.Get(id, password)
}