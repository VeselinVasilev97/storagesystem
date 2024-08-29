package main

import (
	"log"
	"net/http"
	"storage/configuration"
	"storage/middleware"
	"storage/services/categories"
	"storage/services/orders"
	"storage/services/products"

	"github.com/gin-gonic/gin"
)

func main() {
	c := configuration.LoadConfig()

	r := gin.Default()
	r.Use(middleware.LoggingMiddleware)

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "This is the version 1.5.69 - NEW: (endpoint:/api/order GET)")
	})

	apiGroup := r.Group("/api")

	// Products
	apiGroup.GET("/get-products", products.HandlerGetAllProducts(c))
	apiGroup.GET("/get-product", products.HandlerGetProductById(c))
	apiGroup.GET("/get-categories", categories.HandlerGetAllCategories(c))

	// Orders
	apiGroup.POST("/order", orders.HandlerCreateOrder(c))
	apiGroup.GET("/get-order", orders.HandlerGetOrderById(c))
	apiGroup.GET("/orders", orders.HandlerGetAllOrders(c))

	if err := r.Run(":" + c.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
