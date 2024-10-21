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
	"storage/services/suppliers"
	"storage/services/users"
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

		// Routes requiring authentication
		protected := apiGroup.Group("/")
		protected.Use(middleware.AuthMiddleware())

		// Products routes
		{
			productsGroup := protected.Group("/products")

			productsGroup.GET("/get-all", products.HandlerGetAllProducts(d))
			productsGroup.GET("/get-one", products.HandlerGetProductById(d))
		}

		// Categories routes
		{
			categoriesGroup := protected.Group("/categories")

			categoriesGroup.GET("/get-all", categories.HandlerGetAllCategories(d))
			categoriesGroup.GET("/get-one", categories.HandlerGetCategoryById(d))
		}

		// Suppliers routes
		{
			suppliersGroup := protected.Group("/suppliers")

			suppliersGroup.GET("/get-all", suppliers.HandlerGetAllSuppliers(d))
			suppliersGroup.GET("/get-one", suppliers.HandlerGetSupplierById(d))
		}

		// Orders routes
		{
			ordersGroup := protected.Group("/orders")

			ordersGroup.POST("/create", orders.HandlerCreateOrder(d))
			ordersGroup.GET("/get-one", orders.HandlerGetOrderById(d))
			ordersGroup.GET("/get-all", orders.HandlerGetAllOrders(d))

		}
		// Users route
		protected.GET("/users", users.HandlerGetAllUsers(d))
	}

	return r

}
