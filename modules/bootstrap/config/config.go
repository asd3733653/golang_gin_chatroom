package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT      string
	LOG_LEVEL string
}

func ConfigInit() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("loding .env file fail")
	}

	config := *&Config{}
	config.PORT = os.Getenv("PORT")
	config.LOG_LEVEL = os.Getenv("LOG_LEVEL")
	return config
}
