package models

import (
	"time"
)

type Customer struct {
	CustomerID  uint      `json:"customer_id" gorm:"column:customer_id;primary_key;autoIncrement"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Order       []Order
}
