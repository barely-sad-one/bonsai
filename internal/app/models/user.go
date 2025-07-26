package models

import (
	"gorm.io/gorm"
	"gorm.io/datatypes"
	"github.com/google/uuid"
	"time"
)

type Users struct {
	gorm.Model
	ID       	 uuid.UUID 				  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email    	 string             `gorm:"uniqueIndex"`
	Password   string                        
	Name       string                        
	attributes datatypes.JSONMap             
	createdAt  time.Time                     
}
