package persistence

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	db *gorm.DB
}

func Init(dsn string) (*PostgresConfig, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Todo{})
	db.AutoMigrate(&Item{})
	return &PostgresConfig{db: db}, nil
}
