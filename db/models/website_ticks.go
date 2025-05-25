package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WebsiteTick struct {
	gorm.Model
	ID          string        `json:"id" gorm:"primaryKey;type:uuid;"`
	WebsiteID   string        `json:"website_id" gorm:"type:uuid;"`
	Timestamp   time.Time     `json:"timestamp"`
	ValidatorID string        `json:"validator_id" gorm:"type:uuid;"`
	Status      WebsiteStatus `json:"status" gorm:"type:jsonb"`
	Latency     float64       `json:"latency"`
	Website     Website       `json:"website"`
	Validator   Validator     `json:"validator"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (tick *WebsiteTick) BeforeCreate(tx *gorm.DB) error {
	tick.ID = uuid.New().String()
	return nil
}

type WebsiteStatus struct {
	Good int `json:"good"`
	Bad  int `json:"bad"`
}

// Value implements the driver.Valuer interface
func (ws WebsiteStatus) Value() (driver.Value, error) {
	return json.Marshal(ws)
}

// Scan implements the sql.Scanner interface
func (ws *WebsiteStatus) Scan(value interface{}) error {
	if value == nil {
		*ws = WebsiteStatus{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, ws)
}
