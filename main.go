package main

import (
	"log"
	"net/http"
	"storage/configuration"
	"storage/middleware"
	"storage/services/categories"
	"storage/services/orders"
	"storage/services/products"
	"storage/services/suppliers"
	"storage/services/users"

	"github.com/gin-gonic/gin"
)

func main() {
	c := configuration.LoadConfig()

	r := gin.Default()
	r.Use(middleware.LoggingMiddleware)

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "This is the version 1.5.8 - updates: New endpoints for users and suppliers")
	})

	apiGroup := r.Group("/api")

	// Products
	apiGroup.GET("/get-products", products.HandlerGetAllProducts(c))
	apiGroup.GET("/get-product", products.HandlerGetProductById(c))

	// Categories
	apiGroup.GET("/categories", categories.HandlerGetAllCategories(c))
	apiGroup.GET("/category", categories.HandlerGetAllCategories(c))

	// Suppliers
	apiGroup.GET("/suppliers", suppliers.HandlerGetAllSuppliers(c))
	apiGroup.GET("/get-supplier", suppliers.HandlerGetSupplierById(c))

	// Orders
	apiGroup.POST("/order", orders.HandlerCreateOrder(c))
	apiGroup.GET("/get-order", orders.HandlerGetOrderById(c))
	apiGroup.GET("/orders", orders.HandlerGetAllOrders(c))

	// Users
	apiGroup.GET("/users", users.HandlerGetAllUsers(c))

	if err := r.Run(":" + c.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
