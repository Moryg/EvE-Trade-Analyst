package shared

import (
	"time"
)

type NextHref struct {
	Url string `json:"href"`
}

type Order struct {
	VInitial  uint64    `json:"volume_total"`
	Volume    uint64    `json:"volume_remain"`
	Created   time.Time `json:"issued"`
	Id        uint64    `json:"order_id"`
	ItemId    uint64    `json:"type_id"`
	StationId uint64    `json:"location_id"`
	Buy       bool      `json:"is_buy_order"`
	Price     float64   `json:"price"`
	Duration  int       `json:"duration"`
}

func (this *Order) ExpiryDate() string {
	return this.Created.AddDate(0, 0, this.Duration).Format(DB_DATETIME)
}
