package routes

import (
	"app/POS/controllers"
	middlewares "app/POS/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	Api := r.Group("/api")
	{
		// auth
		Api.GET("/", controllers.GetApi)
		Api.POST("/login", controllers.Login)
		Api.POST("/signup", controllers.SignUp)

		// products
		Inventory := Api.Group("/inventory")
		Inventory.Use(middlewares.Authorization())
		{
			Inventory.GET("/", controllers.GetInventories)
			Inventory.GET("/:id", controllers.GetInventory)
			Inventory.POST("/", controllers.CreateInventory)
			Inventory.DELETE("/:id", controllers.DeleteInventory)
			Inventory.PATCH("/:id", controllers.UpdateInventory)
		}

		// customers
		Customer := Api.Group("/customer")
		Customer.Use(middlewares.Authorization())
		{
			Customer.GET("/", controllers.GetCustomers)
			Customer.GET("/:id", controllers.GetCustomer)
			Customer.POST("/", controllers.CreateCustomer)
			Customer.DELETE("/:id", controllers.DeleteCustomer)
			Customer.PATCH("/:id", controllers.UpdateCustomer)
		}

		// orders
		Order := Api.Group("/order")
		Order.Use(middlewares.Authorization())
		{
			Order.GET("/", controllers.GetOrders)
			Order.GET("/:id", controllers.GetCustomer)
			Order.POST("/", controllers.CreateOrder)
			Order.DELETE("/:id", controllers.DeleteCustomer)
			Order.PATCH("/:id", controllers.UpdateCustomer)
		}

	}

	return r
}
