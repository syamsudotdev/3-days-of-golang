package router

import (
	// "fmt"

	"github.com/gin-gonic/gin"

	"ijahinventory/http/controller/productcontroller"
)

func SetRoutes(engine *gin.Engine) {
	engine.GET("/", Root)
	engine.POST("/products", productcontroller.Store)
}

func Root(c *gin.Context) {
	c.JSON(200, "Welcome !")
}
