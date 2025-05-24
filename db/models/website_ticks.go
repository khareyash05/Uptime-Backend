package models

import (
	"time"

	"gorm.io/gorm"
)

type WebsiteTick struct {
	gorm.Model
	ID          string    `json:"id"`
	WebsiteID   string    `json:"website_id"`
	Timestamp   time.Time `json:"timestamp"`
	ValidatorID string    `json:"validator_id"`
	Status      string    `json:"status"`
	Latency     int       `json:"latency"`
}
