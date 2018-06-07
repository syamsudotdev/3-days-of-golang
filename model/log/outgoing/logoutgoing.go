package outgoing

import (
	"time"
)

type LogOutgoing struct {
	ID            int       `json:"id"`
	Timestamp     time.Time `json:"timestamp"`
	ProducId      int       `json:"-"`
	TotalPrice    int       `json:"total_price"`
	CountOutgoing int       `json:"count_outgoing"`
	Note          string    `json:"note"`
}
