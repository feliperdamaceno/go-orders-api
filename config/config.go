package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Host string
	Port string
}

func NewConfig() (*AppConfig, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	config := &AppConfig{
		Host: getEnv("HOST", "localhost"),
		Port: getEnv("PORT", "8000"),
	}

	return config, nil
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
