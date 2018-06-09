package request

import (
	"ijahinventory/data/model/product"
)

type ProductIngoing struct {
	CountOrder    int             `json:"count_order"`
	ReceiptNumber string          `json:"receipt_number"`
	Note          string          `json:"note"`
	Product       product.Product `json:"product"`
}
