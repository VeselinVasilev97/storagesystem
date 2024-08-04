package product

import (
	"net/http"
	"storage/configuration"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandlerGetAllProducts(conf *configuration.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := RepoGetAllProducts(conf.Db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}

}

func HandlerGetProductById(conf *configuration.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		productIdStr := c.Query("id")
		if productIdStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id query parameter is required"})
			return
		}

		productId, err := strconv.ParseInt(productIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id must be a number: " + err.Error()})
			return
		}

		product, err := RepoGetProductById(conf.Db, productId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product: " + err.Error()})
			return
		}

		if product == (Product{}) {
			c.Status(http.StatusNoContent)
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

// func HandlerGetProductByIdDetailed(conf *configuration.Config) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		productIdStr := c.Query("id")
// 		if productIdStr == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "id query parameter is required"})
// 			return
// 		}

// 		productId, err := strconv.ParseInt(productIdStr, 10, 64)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Id must be a number: " + err.Error()})
// 			return
// 		}

// 		product, err := RepoGetProductByIdDetailed(conf.Db, productId)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product: " + err.Error()})
// 			return
// 		}

// 		if product == (Product{}) {
// 			c.Status(http.StatusNoContent)
// 			return
// 		}

// 		c.JSON(http.StatusOK, product)
// 	}
// }
