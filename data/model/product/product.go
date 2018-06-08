package product

type Product struct {
	ID         int    `gorm:"AUTO_INCREMENT" json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Sku        string `gorm:"unique" json:"sku"`
	StockCount int    `json:"stock_count"`
}
