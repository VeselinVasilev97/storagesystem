package main

import (
	"log"
	"net/http"
	"storage/configuration"
	"storage/middleware"
	login "storage/services/Login"
	register "storage/services/Register"
	"storage/services/categories"
	"storage/services/orders"
	"storage/services/products"
	"storage/services/suppliers"
	"storage/services/users"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load configuration
	c := configuration.LoadConfig()

	// Initialize Gin router
	r := gin.Default()
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.CORSandCSP())

	// Define the version endpoint
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "This is version 2.1 - new login form")
	})

	// API route group
	apiGroup := r.Group("/api")
	{
		// Public routes
		apiGroup.POST("/login", login.LoginHandler(c)) // Call the function with the configuration and pass the result
		// Register route
		apiGroup.POST("/register", register.RegisterHandler(c))
		// Routes requiring authentication
		protected := apiGroup.Group("/")
		protected.Use(middleware.AuthMiddleware())

		// Products routes
		protected.GET("/get-products", products.HandlerGetAllProducts(c))
		protected.GET("/get-product", products.HandlerGetProductById(c))

		// Categories routes
		protected.GET("/categories", categories.HandlerGetAllCategories(c))
		protected.GET("/category", categories.HandlerGetCategoryById(c))

		// Suppliers routes
		protected.GET("/suppliers", suppliers.HandlerGetAllSuppliers(c))
		protected.GET("/get-supplier", suppliers.HandlerGetSupplierById(c))

		// Orders routes
		protected.POST("/order", orders.HandlerCreateOrder(c))
		protected.GET("/get-order", orders.HandlerGetOrderById(c))
		protected.GET("/orders", orders.HandlerGetAllOrders(c))
		protected.GET("/orders-today", orders.HandlerGetTodayOrders(c))

		// Users route
		protected.GET("/users", users.HandlerGetAllUsers(c))
	}

	if err := r.Run(":" + c.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
