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
		c.String(http.StatusOK, "This is the version 1.5.2 - PIPEline successfuly set.")
	})

	apiGroup := r.Group("/api")
	apiGroup.GET("/get-products", products.HandlerGetAllProducts(c))
	apiGroup.GET("/get-product", products.HandlerGetProductById(c))
	// apiGroup.GET("/get-product-detailed", product.HandlerGetProductByIdDetailed(c))
	apiGroup.GET("/get-categories", categories.HandlerGetAllCategories(c))
	apiGroup.POST("/order", orders.HandlerCreateOrder(c))
	apiGroup.GET("/get-order", orders.HandlerGetOrderById(c))

	if err := r.Run(":" + c.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
