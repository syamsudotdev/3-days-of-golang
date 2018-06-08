package http

import (
	"ijahinventory/http/router"

	"github.com/gin-gonic/gin"
)

func Start() {
	ginEngine := gin.Default()
	router.SetRoutes(ginEngine)
	ginEngine.Run(":8080")
}
