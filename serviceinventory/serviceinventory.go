package serviceinventory

import (
	"github.com/jinzhu/gorm"

	"ijahinventory/database"
	// "ijahinventory/model/log/ingoing"
	// "ijahinventory/model/log/outgoing"
	"ijahinventory/model/product"
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

func StoreProduct(item product.Product) (*product.Product, error) {
	openDb()
	defer closeDb()

	db.Create(&item)
	if err := db.First(&item, &item.ID).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func GetBySku(sku string) (*product.Product, error) {
	openDb()
	defer closeDb()

	return nil, nil
}
