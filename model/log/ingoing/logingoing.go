package ingoing

import (
	"github.com/jinzhu/gorm"

	"time"
)

type LogIngoing struct {
	gorm.Model
	ID            int       `gorm:"AUTO_INCREMENT" json:"id"`
	Timestamp     time.Time `json:"timestamp"`
	CountOrder    int       `json:"count_order"`
	CountReceived int       `json:"count_received"`
	ProductId     int       `json:"-"`
	Note          string    `json:"note"`
	ReceiptNumber string    `json:"receipt_number"`
	TotalPrice    int       `json:"total_price"`
}
