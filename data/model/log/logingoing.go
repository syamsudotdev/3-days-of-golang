package log

import (
	"ijahinventory/data/model/product"

	"time"
)

type LogIngoing struct {
	ID            int             `gorm:"AUTO_INCREMENT" json:"id"`
	Timestamp     time.Time       `json:"timestamp"`
	CountOrder    int             `json:"count_order"`
	CountReceived int             `json:"count_received"`
	ProductId     int             `json:"-"`
	Product       product.Product `gorm:"foreignkey:ProductId" json:"product"`
	Note          string          `json:"note"`
	ReceiptNumber string          `json:"receipt_number"`
	TotalPrice    int             `json:"total_price"`
}

func (LogIngoing) TableName() string {
	return "log_ingoing"
}
