package config

import (
	"log"
	"os"
)

type AppConfig struct {
	DBUrl string
}

func InitConfig() *AppConfig {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("❌ DATABASE_URL not set in environment")
	}

	return &AppConfig{
		DBUrl: dbURL,
	}
}
