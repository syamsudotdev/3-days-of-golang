package product

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Sku   string `json:"sku"`
	Stock int    `json:"stock"`
}
