package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"time"
)

type Resources struct {
	ID uuid.UUID
	Type string
	OwnerID uuid.UUID
	Attributes datatypes.JSONMap
	CreatedAt time.Time
}
