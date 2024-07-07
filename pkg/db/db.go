package db

import (
	"fmt"
	"tracking_service/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(config configs.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
