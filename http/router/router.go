package router

import (
	// "fmt"

	"github.com/gin-gonic/gin"

	"ijahinventory/http/controller/productcontroller"
)

func SetRoutes(ginEngine *gin.Engine) {
	ginEngine.GET("/", Root)
	ginEngine.POST("/products", productcontroller.Store)
	ginEngine.PUT("/products/out", productcontroller.LogOutgoing)
}

func Root(c *gin.Context) {
	c.JSON(200, "Welcome !")
}
