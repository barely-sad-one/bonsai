package database

import (
	"log"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/barely-sad-one/bonsai/internal/app/config"
)

func GetPostgresDB() *gorm.DB {
	var err error
	var db *gorm.DB

	db, err = gorm.Open(postgres.Open(config.DatabaseConnectionString()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from GROM: %v", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db
}
