package database

import (
	"app/POS/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() {

	u := os.Getenv("DATABASE_USER")
	p := os.Getenv("DATABSE_PASSWORD")
	h := os.Getenv("DATABASE_HOST")
	n := os.Getenv("DATABASE_NAME")
	q := "charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", u, p, h, n, q)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.User{}, &models.Inventory{}, &models.Customer{})

	if err != nil {
		panic("Could not open database connection")
	}

	DB = db
}
