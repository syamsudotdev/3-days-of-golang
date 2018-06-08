package response

import (
	"time"
)

type ResponseLogOutgoing struct {
	Timestamp     time.Time `json:"timestamp"`
	Sku           string    `json:"sku"`
	ProductName   string    `json:"product_name"`
	CountOutgoing int       `json:"count_outgoing"`
	Price         int       `json:"price"`
	TotalPrice    int       `json:"total_price"`
	Note          string    `json:"note"`
}
