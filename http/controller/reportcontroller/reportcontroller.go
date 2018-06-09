package reportcontroller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"ijahinventory/data/service/servicereport"
	"ijahinventory/http/request"
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

func ReportSales(c *gin.Context) {
	var request request.ReportSales
	c.Bind(&request)
	layout := "2006-01-02"
	startTime, err := time.Parse(layout, request.Start)
	if err != nil {
		responsehandler.HandleFile(c, err, path)
	}
	endTime, err := time.Parse(layout, request.End)

	fmt.Println(request.Start)
	fmt.Println(request.End)
	fmt.Println(startTime.Format("2006-01-02"))
	fmt.Println(endTime.Format("2006-01-02"))

	if err != nil {
		responsehandler.HandleFile(c, err, path)
	}
	err = servicereport.GenerateCsvSales(path, startTime, endTime)
	responsehandler.HandleFile(c, err, path)
}
