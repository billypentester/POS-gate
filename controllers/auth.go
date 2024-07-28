package controllers

import (
	"app/POS/database"
	"app/POS/models"
	utils "app/POS/utils/jwt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	var req models.User

	if err := c.BindJSON(&req); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid Request",
		})
		return
	}

	if err := database.DB.Select("email").Where("email = ?", req.Email).First(&req).Error; err != nil {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

		if err != nil {
			log.Printf("failed to hash password: %v", err)
			c.JSON(500, gin.H{
				"status":  false,
				"message": "internal server error",
			})
		}

		if err := database.DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", req.Email, hashedPassword).Error; err != nil {
			log.Printf("failed to insert user into database: %v", err)
			c.JSON(500, gin.H{
				"status":  false,
				"message": "internal server error",
			})
		}

		token, err := utils.CreateToken(req.Email)

		if err := database.DB.Exec("UPDATE users SET token = ? WHERE email = ?", token, req.Email).Error; err != nil {
			log.Printf("failed to save token in database: %v", err)
			c.JSON(500, gin.H{
				"status":  false,
				"message": "internal server error",
			})
		}

		if err != nil {
			log.Printf("failed to create token: %v", err)
			c.JSON(500, gin.H{
				"status":  false,
				"message": "internal server error",
			})
		}

		c.JSON(200, gin.H{
			"status":  true,
			"message": "Account successfully created",
			"token":   token,
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "User already exists",
		})
	}

}

func Login(c *gin.Context) {

	var req models.User

	if err := c.BindJSON(&req); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid Request",
		})
		return
	}

	var user models.User

	if err := database.DB.Select("email", "password").Where("email = ?", req.Email).First(&user).Error; err != nil {
		log.Printf("failed to get user by email from database: %v", err)
		c.JSON(500, gin.H{
			"status":  false,
			"message": "User not found",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Incorrect password",
		})
		return
	}

	token, err := utils.CreateToken(req.Email)

	if err != nil {
		log.Printf("failed to create token: %v", err)
		c.JSON(500, gin.H{
			"status":  false,
			"message": "internal server error",
		})
		return
	}

	if err := database.DB.Exec("UPDATE users SET token = ? WHERE email = ?", token, req.Email).Error; err != nil {
		log.Printf("failed to save token in database: %v", err)
		c.JSON(500, gin.H{
			"status":  false,
			"message": "internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Login successful",
		"token":   token,
	})

}
