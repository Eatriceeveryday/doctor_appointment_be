package config

import (
	"github.com/joho/godotenv"
	"os"
)

var Cfg *Config

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string

	AccessKey string
}

func LoadConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	Cfg = &Config{
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		AccessKey:  os.Getenv("ACCESS_KEY"),
	}
	return nil
}
