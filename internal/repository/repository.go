package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"pastebin/internal/models"
)

type Repository struct {
	Paste Paste
}

func NewRepository(postgres *sqlx.DB, redis *redis.Client, amazon *AmazonDB) *Repository {
	return &Repository{
		Paste: NewPasteRepository(postgres, redis, amazon),
	}
}

type Paste interface {
	Create(userPaste models.UserPaste) (string, error)
	Get(id string, password string) (models.UserPaste, error)
	Delete(id string) error
}
