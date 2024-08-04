package categories

import (
	"net/http"
	"storage/configuration"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandlerGetAllCategories(conf *configuration.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := RepoGetAllCategories(conf.Db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, categories)
	}
}

func HandlerGerCategoryById(conf *configuration.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryIdStr := c.Query("id")
		categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retreive category: " + err.Error()})
			return
		}

		category, err := RepoGetCategoryById(conf.Db, categoryId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retreive category by id: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, category)
	}
}
