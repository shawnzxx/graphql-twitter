package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

type database struct {
	URL string
}

type Config struct {
	Database database
}

func LoadEnv(fileName string) {
	re := regexp.MustCompile(`^(.*` + "twitter" + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	log.Printf("rootPath is: %s", rootPath)

	err := godotenv.Load(string(rootPath) + `/` + fileName)
	if err != nil {
		godotenv.Load()
	}
}

func New() *Config {
	return &Config{
		Database: database{
			URL: os.Getenv("DATABASE_URL"),
		},
	}
}
