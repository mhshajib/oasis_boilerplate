package config

import (
	"github.com/spf13/viper"
)

// Init load configurations from config.yml file
func Init(cfgFile string) error {
	viper.SetEnvPrefix("deals")
	viper.BindEnv("env")
	viper.SetConfigName("config") // name of config file (without extension)
	if cfgFile != "" {
		viper.SetConfigName(cfgFile)
	}
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		return err
	}
	initConfig()
	return nil
}

// initConfig laod all configurations
func initConfig() {
	loadApp()
	loadDatabase()
	loadSmsManager()
	loadStorageManager()
	loadRedis()
	loadCORS()
	loadJWT()
}
