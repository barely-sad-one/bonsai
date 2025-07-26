package models

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID       uuid.UUID
	RefreshToken string    `gorm:"uniqueIndex"`
	UserAgent    string
	IP           string
	ExpiresAt    time.Time
	CreatedAt    time.Time
}
