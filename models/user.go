package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"->;primaryKey"`
	Email     string    `json:"email" gorm:"index:,unique"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
