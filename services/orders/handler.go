package orders

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"storage/configuration"
)

func HandlerCreateOrder(conf *configuration.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newOrder NewOrder
		if err := c.BindJSON(&newOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		orderId, err := RepoCreateNewOrder(conf.Db, newOrder)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "id": orderId})
	}

}
