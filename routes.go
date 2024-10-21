package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"storage/configuration"
	"storage/middleware"
	"storage/services/categories"
	login "storage/services/login"
	"storage/services/orders"
	"storage/services/products"
	register "storage/services/register"
	"storage/services/suppliers"
	"storage/services/user"
)

func Routes(d *configuration.Dependencies) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggingMiddleware)

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "This is version 2.0 - updates:LOGIN authentication JWT added.")
	})

	apiGroup := r.Group("/api")
	{
		// Public routes
		apiGroup.POST("/login", login.LoginHandler(d))
		// Register route
		apiGroup.POST("/register", register.RegisterHandler(d))

		// Routes requiring authentication
		protected := apiGroup.Group("/")
		protected.Use(middleware.AuthMiddleware())

		// Products routes
		{
			productsGroup := protected.Group("/products")

			productsGroup.GET("/", products.HandlerGetAllProducts(d))
			productsGroup.GET("/:id", products.HandlerGetProductById(d))
		}

		// Categories routes
		{
			categoriesGroup := protected.Group("/categories")

			categoriesGroup.GET("/", categories.HandlerGetAllCategories(d))
			categoriesGroup.GET("/:id", categories.HandlerGetCategoryById(d))
		}

		// Suppliers routes
		{
			suppliersGroup := protected.Group("/suppliers")

			suppliersGroup.GET("/", suppliers.HandlerGetAllSuppliers(d))
			suppliersGroup.GET("/:id", suppliers.HandlerGetSupplierById(d))
		}

		// Orders routes
		{
			ordersGroup := protected.Group("/orders")

			ordersGroup.POST("/", orders.HandlerCreateOrder(d))
			ordersGroup.GET("/:id", orders.HandlerGetOrderById(d))
			ordersGroup.GET("/", orders.HandlerGetAllOrders(d))

		}
		// Users route
		protected.GET("/users", user.HandlerGetAllUsers(d))
	}

	return r

}
