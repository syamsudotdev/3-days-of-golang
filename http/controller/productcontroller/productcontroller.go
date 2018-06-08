package productcontroller

import (
	"github.com/gin-gonic/gin"

	"ijahinventory/data/serviceproduct"
	"ijahinventory/http/request"
	"ijahinventory/http/responsehandler"
)

func Store(c *gin.Context) {
	var request request.RequestProductIngoing
	c.Bind(&request)

	result, err := serviceproduct.StoreProduct(
		request.Product,
		request.CountOrder,
		request.ReceiptNumber,
		request.Note,
	)
	responsehandler.Handle(c, err, result)
}

func LogOutgoing(c *gin.Context) {
	var request request.RequestProductOutgoing
	c.Bind(&request)

	result, err := serviceproduct.LogOutgoing(
		request.Sku,
		request.CountOutgoing,
		request.Note,
	)
	responsehandler.Handle(c, err, result)
}
