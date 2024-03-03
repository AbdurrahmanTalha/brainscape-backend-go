package db

import (
	"fmt"
	"time"

	"github.com/AbdurrahmanTalha/brainscape-backend-go/config"
	models "github.com/AbdurrahmanTalha/brainscape-backend-go/data/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func SetupDB(cfg *config.Config) error {
	var err error
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DatabaseName, cfg.Postgres.SSLMode)

	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})

	if err != nil {
		return err
	}

	sqlDB, _ := dbClient.DB()
	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	err = autoMigrate()

	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}

func autoMigrate() error {
	err := dbClient.AutoMigrate(&models.User{})

	return err
}
