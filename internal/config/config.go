package config

import (
	"os"

	"github.com/joho/godotenv"
)

type PortConfig struct {
	Port string
}

func LoadPort() PortConfig {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return PortConfig{Port: port}
}
