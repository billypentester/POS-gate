package controllers

import (
	"app/POS/database"
	"app/POS/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInventories(c *gin.Context) {

	var inventory []models.Inventory

	if err := database.DB.Find(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to fetch inventory data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   inventory,
	})

}

func GetInventory(c *gin.Context) {

	var inventory models.Inventory
	id := c.Param("id")

	if err := database.DB.First(&inventory, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to fetch inventory data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   inventory,
	})

}

func CreateInventory(c *gin.Context) {

	var inventory models.Inventory

	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid request payload",
		})
		return
	}

	if err := database.DB.Create(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to create inventory",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   inventory,
	})

}

func DeleteInventory(c *gin.Context) {

	id := c.Param("id")
	var inventory models.Inventory

	if err := database.DB.First(&inventory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "Inventory not found",
		})
		return
	}

	if err := database.DB.Delete(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to delete inventory",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Inventory deleted successfully",
	})

}

func UpdateInventory(c *gin.Context) {

	var updatedInventory models.Inventory
	id := c.Param("id")

	if err := c.ShouldBindJSON(&updatedInventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid request payload",
		})
		return
	}

	var inventory models.Inventory
	if err := database.DB.First(&inventory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "Inventory not found",
		})
		return
	}

	if err := database.DB.Model(&inventory).Updates(updatedInventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to update inventory",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   inventory,
	})

}
