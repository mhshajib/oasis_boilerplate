package config

import (
	"github.com/spf13/viper"
)

// CORSCfg represents jwt config
type CORSCfg struct {
	Origins []string `json:"origins"`
	MaxAge  int      `json:"max_age"`
}

var cors CORSCfg

// CORS contains cors configurations
func CORS() *CORSCfg {
	return &cors
}

func loadCORS() {
	cors = CORSCfg{
		Origins: viper.GetStringSlice("cors.origins"),
		MaxAge:  viper.GetInt("cors.max_age"),
	}
}
