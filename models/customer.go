package models

import (
	"time"
)

type Customer struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Category    string    `json:"category`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
