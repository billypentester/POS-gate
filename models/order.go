package models

import (
	"time"
)

type Order struct {
	OrderID    uint      `json:"order_id" gorm:"column:order_id;primary_key;autoIncrement"`
	Price      float64   `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	OrderItem  []OrderItem
	CustomerID uint
	Customer   Customer
}
