package productcontroller

import (
	"github.com/gin-gonic/gin"

	"ijahinventory/data/serviceproduct"
	"ijahinventory/http/request"
	"ijahinventory/http/responsehandler"
)

func Store(c *gin.Context) {
	var request request.ProductIngoing
	c.Bind(&request)

	result, err := serviceproduct.StoreProduct(
		request.Product,
		request.CountOrder,
		request.ReceiptNumber,
		request.Note,
	)
	responsehandler.Handle(c, err, result)
}
