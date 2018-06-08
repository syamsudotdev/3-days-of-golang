package serviceproduct

import (
	// "fmt"
	"errors"
	"time"

	"github.com/jinzhu/gorm"

	"ijahinventory/data/database"
	"ijahinventory/data/model/log"
	// "ijahinventory/model/log/outgoing"
	"ijahinventory/data/model/product"
)

var (
	db  *gorm.DB
	err error
)

func openDb() {
	db, err = database.OpenDb()
}

func closeDb() {
	db.Close()
}

//store new product or update existing row and log it
//if sku not found and other product's fields are empty, return error
func StoreProduct(item product.Product,
	countOrder int,
	receiptNumber string,
	note string) (*product.Product, error) {
	openDb()
	defer closeDb()

	stockCount := item.StockCount
	checkRow := db.First(&item, product.Product{Sku: item.Sku})
	if !checkRow.RecordNotFound() {
		item.StockCount = stockCount
	} else {
		if item.Name == "" || item.Price == 0 || item.StockCount == 0 {
			return nil,
				errors.New("name, price, or stock_count cannot be empty")
		}
		if err := db.Create(&item).Error; err != nil {
			return nil, err
		}
	}

	logIngoing := log.LogIngoing{
		Timestamp:     time.Now(),
		CountOrder:    countOrder,
		CountReceived: item.StockCount,
		Product:       item,
		Note:          note,
		ReceiptNumber: receiptNumber,
		TotalPrice:    item.Price * item.StockCount,
	}
	if err := db.Create(&logIngoing).Error; err != nil {
		return nil, err
	}

	return &item, nil
}
