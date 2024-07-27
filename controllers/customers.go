package controllers

import (
	"app/POS/database"
	"app/POS/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {

	var customer []models.Customer

	if err := database.DB.Find(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to fetch customer data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   customer,
	})

}

func GetCustomer(c *gin.Context) {

	var customer models.Customer
	id := c.Param("id")

	if err := database.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to fetch customer data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   customer,
	})

}

func CreateCustomer(c *gin.Context) {

	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid request payload",
		})
		return
	}

	if err := database.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to create customer",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   customer,
	})

}

func DeleteCustomer(c *gin.Context) {

	id := c.Param("id")
	var customer models.Customer

	if err := database.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "Customer not found",
		})
		return
	}

	if err := database.DB.Delete(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to delete customer",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Customer deleted successfully",
	})

}

func UpdateCustomer(c *gin.Context) {

	var updatedCustomer models.Customer
	id := c.Param("id")

	if err := c.ShouldBindJSON(&updatedCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid request payload",
		})
		return
	}

	var customer models.Customer
	if err := database.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "Customer not found",
		})
		return
	}

	if err := database.DB.Model(&customer).Updates(updatedCustomer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to update customer",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   customer,
	})

}
