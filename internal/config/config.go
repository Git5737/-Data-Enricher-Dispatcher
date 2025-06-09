package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiA string
	ApiB string
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default env")
	}

	return Config{
		ApiA: os.Getenv("API_A_URL"),
		ApiB: os.Getenv("API_B_URL"),
	}
}
