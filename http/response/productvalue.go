package response

import (
	"time"
)

type ReportProductValueItem struct {
	Sku             string `json:"sku"`
	ProductName     string `json:"product_name"`
	Quantity        int    `json:"quantity"`
	AverageBuyPrice int    `json:"average_buy_price"`
	TotalPrice      int    `json:"total_price"`
}

type ReportProductValue struct {
	Timestamp     time.Time                `json:"print_time"`
	SkuCount      int                      `json:"sku_count"`
	TotalQuantity int                      `json:"total_quantity"`
	TotalPrice    int                      `json:"total_price"`
	Item          []ReportProductValueItem `json:"item"`
}
