package reportcontroller

import (
	"github.com/gin-gonic/gin"

	"ijahinventory/data/service/servicereport"
	"ijahinventory/http/responsehandler"
)

var path = "result.csv"

func ReportInventoryCount(c *gin.Context) {
	err := servicereport.GenerateCsvProductCount(path)
	responsehandler.HandleFile(c, err, path)
}

func ReportProductIngoing(c *gin.Context) {
	err := servicereport.GenerateCsvProductIngoing(path)
	responsehandler.HandleFile(c, err, path)
}

func ReportProductOutgoing(c *gin.Context) {
	err := servicereport.GenerateCsvProductOutgoing(path)
	responsehandler.HandleFile(c, err, path)
}

func ReportProductValue(c *gin.Context) {
	err := servicereport.GenerateCsvProductValue(path)
	responsehandler.HandleFile(c, err, path)
}
