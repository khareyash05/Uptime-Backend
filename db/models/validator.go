package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Validator struct {
	gorm.Model
	ID           string        `json:"id" gorm:"primaryKey;type:uuid;"`
	PublicKey    string        `json:"public_key"`
	IP           string        `json:"ip"`
	WebsiteTicks []WebsiteTick `json:"website_ticks"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (validator *Validator) BeforeCreate(tx *gorm.DB) error {
	validator.ID = uuid.New().String()
	return nil
}
