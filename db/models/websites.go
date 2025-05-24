package models

import "gorm.io/gorm"

type Website struct {
	gorm.Model
	ID        string `json:"id" gorm:"primaryKey"`
	URL       string `json:"url"`
	UserId    string `json:"user_id"`
	WebsiteTicks []WebsiteTick `json:"website_ticks"`
}