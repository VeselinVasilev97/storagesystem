package orders

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"storage/configuration"
	"strconv"
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

func HandlerGetOrderById(conf *configuration.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderIdStr := c.Query("id")
		if orderIdStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id query parameter is required"})
			return
		}

		orderId, err := strconv.ParseInt(orderIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id must be a number: " + err.Error()})
			return
		}

		order, err := RepoGetOrderById(conf.Db, orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order: " + err.Error()})
			return
		}

		//if order == ([]OrderView){
		//	c.Status(http.StatusNoContent)
		//	return
		//}

		c.JSON(http.StatusOK, order)
	}
}