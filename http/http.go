package http

import (
	"github.com/gin-gonic/gin"
)

func StartListen() {
	http := gin.Default()

	http.GET("/", Root)

	http.Run(":8080")
}

func Root(c *gin.Context) {
	c.JSON(200, "Welcome !")
}
