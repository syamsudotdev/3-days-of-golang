package http

import (
	"ijahinventory/model/product"
	"ijahinventory/request/stockbody"
	"ijahinventory/serviceinventory"

	"github.com/gin-gonic/gin"
)

func StartListen() {
	http := gin.Default()

	http.GET("/", Root)
	http.POST("/stock", StockStore)

	http.Run(":8080")
}

func Root(c *gin.Context) {
	c.JSON(200, "Welcome !")
}

func StockStore(c *gin.Context) {
	var request stockbody.Stock
	c.Bind(&request)

	var item product.Product
	item.Name = request.Name
	item.Price = request.Price
	item.Sku = request.Sku
	item.Stock = request.Stock

	result, err := serviceinventory.StoreProduct(item)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, gin.H{"message": "success", "data": result})
}
