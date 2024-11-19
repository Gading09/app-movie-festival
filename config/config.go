package config

import (
	"movie-festival/database"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	Port     int
	Database database.DatabaseConfig
}

func LoadConfig() (config EnvironmentConfig, err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		return
	}

	config = EnvironmentConfig{
		Port: port,
		Database: database.DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Name:     os.Getenv("DB_NAME"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Port:     os.Getenv("DB_PORT"),
		},
	}

	return
}
