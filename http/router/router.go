package router

import (
	// "fmt"

	"github.com/gin-gonic/gin"

	"ijahinventory/http/controller/productcontroller"
	"ijahinventory/http/controller/reportcontroller"
)

func SetRoutes(ginEngine *gin.Engine) {
	ginEngine.GET("/", Root)
	ginEngine.POST("/products", productcontroller.Store)
	ginEngine.PUT("/products/out", productcontroller.LogOutgoing)
	ginEngine.GET(
		"/products/report/stock-count.csv",
		reportcontroller.ReportInventoryCount,
	)
	ginEngine.GET(
		"/products/report/product-ingoing.csv",
		reportcontroller.ReportProductIngoing,
	)
	ginEngine.GET(
		"/products/report/product-outgoing.csv",
		reportcontroller.ReportProductOutgoing,
	)
	ginEngine.GET(
		"/products/report/product-value.csv",
		reportcontroller.ReportProductValue,
	)
	ginEngine.POST(
		"/products/report/sales.csv",
		reportcontroller.ReportSales,
	)
}

func Root(c *gin.Context) {
	c.JSON(200, "Welcome !")
}
