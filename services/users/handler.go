package users

import (
	"net/http"
	"storage/configuration"

	"github.com/gin-gonic/gin"
)

func HandlerGetAllUsers(conf *configuration.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := RepoGetAllUsers(conf.Db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve suppliers: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}
