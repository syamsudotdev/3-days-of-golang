package responsehandler

import (
	"github.com/gin-gonic/gin"
)

func Handle(c *gin.Context, err error, data interface{}) {
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{"message": "success", "data": data})
	}
}
