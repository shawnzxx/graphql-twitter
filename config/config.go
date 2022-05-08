package config

import (
	"github.com/joho/godotenv"
	"os"
)

type database struct {
	URL string
}

type Config struct {
	Database database
}

//New create a new config which have settings for application to use, like DATABASE_URL
func New() *Config {
	godotenv.Load()

	config := &Config{
		Database: database{
			URL: os.Getenv("DATABASE_URL"),
		},
	}
	return config
}
