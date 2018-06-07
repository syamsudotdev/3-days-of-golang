package outgoing

import (
	"github.com/jinzhu/gorm"

	"time"
)

type LogOutgoing struct {
	gorm.Model
	ID            int       `gorm:"AUTO_INCREMENT" json:"id"`
	Timestamp     time.Time `json:"timestamp"`
	ProducId      int       `json:"-"`
	TotalPrice    int       `json:"total_price"`
	CountOutgoing int       `json:"count_outgoing"`
	Note          string    `json:"note"`
}
