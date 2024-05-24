package repository

import (
	"bytes"
	"encoding/base64"
	b64 "encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"io"
	"pastebin/internal/models"
	"strconv"
	"time"
)

const (
	pasteTTL = time.Hour * 6
	pasteTable = "Paste"
)

type PasteRepository struct {
	postgresql_db *sqlx.DB
	redis_db      *redis.Client
	amazon_db     *AmazonDB
}

func NewPasteRepository(postgresql_db *sqlx.DB, redis_db *redis.Client, amazon_db *AmazonDB) *PasteRepository {
	return &PasteRepository{
		postgresql_db: postgresql_db,
		redis_db:      redis_db,
		amazon_db:     amazon_db,
	}
}

func(r *PasteRepository) Create(userPaste models.UserPaste) (string, error) {
	var pasteId int
	createPasteQuery := fmt.Sprintf("INSERT INTO %s (password, expires_at) values ($1, $2) RETURNING id", pasteTable)

	row := r.postgresql_db.QueryRow(createPasteQuery, userPaste.Password, time.Now().Add(pasteTTL))
	err := row.Scan(&pasteId)
	if err != nil {
		return "", err
	}
	stringPasteId := strconv.Itoa(pasteId)

	content := []byte(userPaste.Text)
	contentLength := int64(len(content))
	hash := b64.StdEncoding.EncodeToString([]byte(stringPasteId))
	if err != nil {
		return "", err
	}

	_, err = r.amazon_db.svc.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String("amazon-pastebin"),
		Key:           aws.String(string(hash)),
		Body:          bytes.NewReader(content),
		ContentLength: aws.Int64(contentLength),
		ContentType:   aws.String("text/plain"),
	})

	if err != nil {
		return "", fmt.Errorf("unable to upload to bucket %v\n", err)
	}

	return hash, nil
}

func(r *PasteRepository) Get(id string, password string) (string, error) {
	//check redis
	text, err := r.redis_db.HGet(ctx, id, "text").Result()
	if err == nil {
		storedPassword, _ := r.redis_db.HGet(ctx, id, "password").Result()
		if password != storedPassword && storedPassword != "" {
			return "", fmt.Errorf("incorrect password")
		}
		return text, nil
	}
	var userPaste models.PostgresPaste

	decodedId, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return "", nil
	}

	query := fmt.Sprintf("SELECT password, expires_at FROM %s WHERE id = $1", pasteTable)
	err = r.postgresql_db.Get(&userPaste, query, decodedId)
	if err != nil {
		return "", err
	}

	if time.Now().After(userPaste.ExpiresAt) {
		return "", fmt.Errorf("expired")
	}

	if password != userPaste.Password && userPaste.Password != "" {
		return "", fmt.Errorf("incorrect password")
	}

	result, err := r.amazon_db.svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("amazon-pastebin"),
		Key:    aws.String(id),
	})
	if err != nil {
		return "", fmt.Errorf("error retrieving object from S3: %v", err)
	}

	objectContent, err := io.ReadAll(result.Body)
	if err != nil {
		return "", fmt.Errorf("error reading object content: %v", err)
	}

	err = r.redis_db.HSet(ctx, id, map[string]interface{}{
		"text":     string(objectContent),
		"password": password,
	}).Err()
	if err != nil {
		fmt.Printf("Error setting value in Redis: %v\n", err)
	}

	return string(objectContent), nil
}