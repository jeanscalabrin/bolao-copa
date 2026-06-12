package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL        string
	FootballDataAPIKey string
}

func Load() Config {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env not found")
	}

	return Config{
		DatabaseURL:        os.Getenv("DATABASE_URL"),
		FootballDataAPIKey: os.Getenv("FOOTBALL_DATA_API_KEY"),
	}
}
