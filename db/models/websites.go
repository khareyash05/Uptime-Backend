package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Website struct {
	gorm.Model
	ID           string        `json:"id" gorm:"primaryKey;type:uuid;"`
	URL          string        `json:"url"`
	UserId       string        `json:"user_id"`
	WebsiteTicks []WebsiteTick `json:"website_ticks"`
	Disabled     bool          `gorm:"default:false"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (website *Website) BeforeCreate(tx *gorm.DB) error {
	website.ID = uuid.New().String()
	return nil
}
