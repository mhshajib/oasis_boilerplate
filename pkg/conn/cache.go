package conn

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mhshajib/oasis_boilerplate/pkg/cache"
	"github.com/mhshajib/oasis_boilerplate/pkg/config"
	"github.com/mhshajib/oasis_boilerplate/pkg/log"
)

var defaultCache cache.Cache
var redisClient *redis.Client

// GetRedis return defautl connected redis client
func GetRedis() *redis.Client {
	return redisClient
}

// ConnectCache ...
func ConnectCache(cfg *config.RedisConfig) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})
	defaultCache = cache.NewRedis(rdb)
	redisClient = rdb
	return rdb.Ping(rdb.Context()).Err()
}

// ConnectDefaultCache connect with default configurations
func ConnectDefaultCache() error {
	cfg := config.Redis()
	err := ConnectCache(cfg)
	// run a background process to ping and establish connection
	go func() {
		for {
			if err := defaultCache.Ping(context.Background()); err != nil {
				log.Warn("cache: ping error:", err)
				if err := ConnectCache(cfg); err != nil {
					log.Warn("cache:failed to reconnect:", err)
				}
			}
			time.Sleep(3 * time.Second)
		}
	}()
	return err
}

// DefaultCache return default connected cache
func DefaultCache() cache.Cache {
	return defaultCache
}
