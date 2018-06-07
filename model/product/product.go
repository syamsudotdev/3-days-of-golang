package product

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ID    int    `gorm:"AUTO_INCREMENT" json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Sku   string `json:"sku"`
	Stock int    `json:"stock"`
}
