package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"time"
)

type ABACPolicy struct {
	ID                uuid.UUID           `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name              string
	Action            string
	SubjectFilter     datatypes.JSONMap
	ResourceFilter    datatypes.JSONMap
	Effect            bool
	Priority          int
	Enable            bool
	CreatedAt         time.Time
}
