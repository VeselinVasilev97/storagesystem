package suppliers

import (
	"net/http"
	"storage/configuration"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandlerGetAllSuppliers(conf configuration.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		suppliers, err := RepoGetAllSuppliers(conf.Db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve suppliers: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, suppliers)
	}
}

func HandlerGetSupplierById(conf configuration.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		supplierIdStr := c.Query("id")
		if supplierIdStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id query parameter is required"})
			return
		}
		supplierId, err := strconv.ParseInt(supplierIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id must be a number: " + err.Error()})
			return
		}

		supplier, err := RepoGetSupplierById(conf.Db, supplierId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retreive supplier by id: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, supplier)
	}
}
