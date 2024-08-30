package storage

import (
	"context"
	"time"
)

type STORAGE interface {
	GeneratePresignedUploadURL(ctx context.Context, bucket, key string, expireMinutes time.Duration) (string, string, error)
	CheckFileExists(ctx context.Context, bucket, key string) bool
}
