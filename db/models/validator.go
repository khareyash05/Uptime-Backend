package models

import "gorm.io/gorm"

type Validator struct {
	gorm.Model
	ID        string `json:"id"`
	PublicKey string `json:"public_key"`
	IP        string `json:"ip"`
}