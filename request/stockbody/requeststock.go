package stockbody

type Stock struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Sku   string `json:"sku"`
	Stock int    `json:"stock"`
}
