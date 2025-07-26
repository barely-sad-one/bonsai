package database

import (
	"sync"
	"gorm.io/gorm"
	"github.com/barely-sad-one/bonsai/internal/app/models"
)

type Database struct {
	instance *gorm.DB
	once     sync.Once
}

var DB Database



func dbMigration() {
	DB.instance.AutoMigrate(
		&models.Users{},
		&models.Resources{},
		&models.ABACPolicy{},
		&models.Session{},
	)
}

func InitDB() {
	DB.once.Do(func() {
		DB.instance = GetPostgresDB()
		DB.instance.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
		dbMigration()
	})
}
