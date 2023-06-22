package database

import (
	"fmt"
	"workspace-service/core/config"

	"github.com/iyiola-dev/go-graphql/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(config *config.DatabaseConfig) (*gorm.DB, error) {
	fmt.Printf("%+v\n", config)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Book{})
}
