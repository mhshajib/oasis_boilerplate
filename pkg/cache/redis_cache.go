package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mhshajib/oasis_boilerplate/domain"
)

// NewRedis return a new redis cache
func NewRedis(client *redis.Client) Cache {
	return &Redis{client: client}
}

// Redis represents a concrete redis
type Redis struct {
	client *redis.Client
}

// Ping ping the redis redis if success return nil
func (r *Redis) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

// Set set a key in redis
func (r *Redis) Set(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	return r.client.Set(ctx, key, value, exp).Err()
}

// Get get a key from redis
func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	resStr, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", domain.ErrCacheNotFound
		}
		return "", err
	}
	return resStr, nil
}

// Del remove a key from redis
func (r *Redis) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// IncrBy increment a value by a key in redis
func (r *Redis) IncrBy(ctx context.Context, key string, value int64) error {
	return r.client.IncrBy(ctx, key, value).Err()
}

// DecrBy decrement a value by a key in redis
func (r *Redis) DecrBy(ctx context.Context, key string, value int64) error {
	return r.client.DecrBy(ctx, key, value).Err()
}
