package controllers

import (
	"app/POS/database"
	"app/POS/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {

	var order []models.Order

	if err := database.DB.Preload("OrderItem").Preload("Customer").Find(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to fetch order data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   order,
	})

}

func GetOrder(c *gin.Context) {

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

func CreateOrder(c *gin.Context) {

	type OrderItem struct {
		ProductID uint    `json:"id"`
		Name      string  `json:"name"`
		Price     float64 `json:"price"`
		Quantity  int     `json:"quantity"`
	}

	type Order struct {
		CustomerID uint        `json:"customer_id"`
		Price      float64     `json:"price"`
		Items      []OrderItem `json:"items"`
	}

	var orderData Order

	if err := c.ShouldBindJSON(&orderData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid request payload",
		})
		return
	}

	var order models.Order

	order.Price = orderData.Price
	order.CustomerID = orderData.CustomerID

	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to create create order",
		})
		return
	}

	if err := database.DB.Preload("Customer").First(&order, order.OrderID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to preload customer",
		})
		return
	}

	var items []*models.OrderItem

	for _, item := range orderData.Items {
		items = append(items, &models.OrderItem{
			OrderID:  order.OrderID,
			Name:     item.Name,
			Price:    item.Price,
			Quantity: item.Quantity,
		})
	}

	if err := database.DB.Create(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to create order items",
		})
	}

	responseData := map[string]interface{}{
		"order": order,
		"items": items,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   responseData,
	})

}

func DeleteOrder(c *gin.Context) {

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

func UpdateOrder(c *gin.Context) {

	var updatedInventory map[string]interface{}
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
