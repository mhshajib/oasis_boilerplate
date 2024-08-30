package cache

import (
	"context"
	"time"
)

// Cache represents cache contract
type Cache interface {
	Ping(ctx context.Context) error
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
	IncrBy(ctx context.Context, key string, value int64) error
	DecrBy(ctx context.Context, key string, value int64) error
}
