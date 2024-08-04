package models

import (
	"time"
)

type OrderItem struct {
	OrderItemID uint      `json:"order_item_id" gorm:"column:order_item_id;primary_key;autoIncrement"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	OrderID     uint      `json:"order_id"`
	// Order       Order
}
