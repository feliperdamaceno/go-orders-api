package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	HOST string
	PORT string
}

var Config = AppConfig{}

func LoadConfig() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	Config.HOST = getEnv("HOST", "localhost")
	Config.PORT = getEnv("PORT", "8000")

	return nil
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
