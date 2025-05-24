package models

import "gorm.io/gorm"

type Validator struct {
	gorm.Model
	ID           string        `json:"id" gorm:"primaryKey"`
	PublicKey    string        `json:"public_key"`
	IP           string        `json:"ip"`
	WebsiteTicks []WebsiteTick `json:"website_ticks"`
}
