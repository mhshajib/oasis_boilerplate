package s3

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Provider struct {
	KeyId                      string
	KeySecret                  string
	Region                     string
	Bucket                     string
	Timeout                    time.Duration
	PresignedUrlExpirationMins time.Duration
	client                     *s3.S3
}

func newS3Session(cfg S3Provider) *s3.S3 {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(cfg.Region),
		Credentials: credentials.NewStaticCredentials(cfg.KeyId, cfg.KeySecret, ""),
		//LogLevel:    aws.LogLevel(aws.LogDebugWithSigning | aws.LogDebugWithHTTPBody),
	}))

	return s3.New(sess)
}

func (s S3Provider) GeneratePresignedUploadURL(ctx context.Context, bucket, key string, expireMinutes time.Duration) (string, string, error) {
	s3Sess := newS3Session(s)
	req, _ := s3Sess.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		ACL:    aws.String("public-read"),
	})
	req.SetContext(ctx) // Apply the context to the request.
	presignedUrl, err := req.Presign(expireMinutes)
	if err != nil {
		return "", "", fmt.Errorf("failed to presign the request: %v", err)
	}

	publicUrl := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, s.Region, key)

	return presignedUrl, publicUrl, nil
}

func (s S3Provider) CheckFileExists(ctx context.Context, bucket, key string) bool {
	s3Sess := newS3Session(s)

	_, err := s3Sess.HeadObject((&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}))
	if err != nil {
		return false
	}
	return true
}
