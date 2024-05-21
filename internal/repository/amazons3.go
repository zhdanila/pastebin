package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AmazonConfig struct {
	Region    string
	AccessKey string
	SecretKey string
}

type AmazonDB struct {
	svc *s3.S3
}

func NewAmazonDB(config AmazonConfig) (*AmazonDB, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Region),
		Credentials: credentials.NewStaticCredentials(config.AccessKey, config.SecretKey, ""),
	})
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	_, err = svc.ListBuckets(nil)
	if err != nil {
		return nil, err
	}

	return &AmazonDB{svc: svc}, nil
}