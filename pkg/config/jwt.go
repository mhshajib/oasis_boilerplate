package config

import (
	"time"

	"github.com/spf13/viper"
)

// JWTConfig represents jwt config
type JWTConfig struct {
	ExpirationTime time.Duration `json:"token_expiration_time"`
	Key            []byte        `json:"key"`
}

var jwt JWTConfig

// JWT contains jwt configurations
func JWT() JWTConfig {
	return jwt
}

func loadJWT() {
	jwt = JWTConfig{
		ExpirationTime: viper.GetDuration("jwt.token_expiration_time") * time.Minute,
		Key:            []byte(viper.GetString("jwt.key")),
	}
}
