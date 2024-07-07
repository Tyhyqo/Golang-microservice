package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	LogLevel   string
	DB         DBConfig
	JWTSecret  string
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		ServerPort: os.Getenv("SERVER_PORT"),
		LogLevel:   os.Getenv("LOG_LEVEL"),
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
