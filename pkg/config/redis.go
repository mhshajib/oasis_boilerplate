package config

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/spf13/viper"
)

// RedisConfig holds the redis configuration
type RedisConfig struct {
	Address  string
	Password string
	DB       int
	WorkerDB int
	Prefix   string
}

// URI build the redis uri from the configuration
func (r *RedisConfig) URI() string {
	u := url.URL{
		Scheme: "redis",
		Host:   r.Address,
		Path:   strconv.Itoa(r.DB),
	}
	if r.Password != "" {
		u.User = url.User(r.Password)
	}
	return u.String()
}

var redis *RedisConfig

// Redis returns the default Redis configuration
func Redis() *RedisConfig {
	return redis
}

// loadRedis loads Redis configuration
func loadRedis() {
	redis = &RedisConfig{
		Address:  fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		WorkerDB: viper.GetInt("redis.worker_db"),
		Prefix:   viper.GetString("redis.prefix"),
	}
}
