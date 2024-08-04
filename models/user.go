package models

import (
	"time"
)

type User struct {
	UserID    uint      `json:"user_id" gorm:"primaryKey;autoIncrement"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" gorm:"index:,unique"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
