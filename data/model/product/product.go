package product

type Product struct {
	//prevent id field to show on response
	ID         int    `gorm:"AUTO_INCREMENT" json:"-"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Sku        string `gorm:"unique" json:"sku"`
	StockCount int    `json:"stock_count"`
}
