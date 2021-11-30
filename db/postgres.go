package db

import (
	"khidr/todo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Todo{})
	db.AutoMigrate(&models.Item{})
	return db, nil
}
