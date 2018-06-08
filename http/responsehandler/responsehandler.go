package responsehandler

import (
	"github.com/gin-gonic/gin"
	"os"
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

func HandleFile(c *gin.Context, err error, path string) {
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": err.Error(),
		})
	} else {
		c.File(path)
		defer os.Remove(path)
	}
}
