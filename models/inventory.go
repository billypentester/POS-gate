package models

import (
	"time"
)

type Inventory struct {
	ProductID   uint      `json:"product_id" gorm:"column:product_id;primary_key;autoIncrement"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
