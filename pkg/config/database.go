package config

import "github.com/spf13/viper"

// Database represents database configuration
type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var db Database

// DB contains database configuration
func DB() Database {
	return db
}

func loadDatabase() {
	db = Database{
		Host:     viper.GetString("mongo_db.host"),
		Port:     viper.GetInt("mongo_db.port"),
		Name:     viper.GetString("mongo_db.name"),
		Username: viper.GetString("mongo_db.username"),
		Password: viper.GetString("mongo_db.password"),
	}
}
