package main

import (
	"log"
	"net/http"
	"storage/configuration"
	"storage/middleware"
	"storage/services/categories"
	"storage/services/product"

	"github.com/gin-gonic/gin"
)

func main() {
	c := configuration.LoadConfig()

	r := gin.Default()
	r.Use(middleware.LoggingMiddleware)

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "This is the version 1.0 of the backend app")
	})

	apiGroup := r.Group("/api")
	apiGroup.GET("/get-products", product.HandlerGetAllProducts(c))
	apiGroup.GET("/get-product", product.HandlerGetProductById(c))
	// apiGroup.GET("/get-product-detailed", product.HandlerGetProductByIdDetailed(c))
	apiGroup.GET("/get-categories", categories.HandlerGetAllCategories(c))

	if err := r.Run(":" + c.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
