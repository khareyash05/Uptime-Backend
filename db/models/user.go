package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `json:"id" gorm:"primaryKey;type:uuid;"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
	return nil
}
