/**
 * Common app structures
 */
package shared

type NextHref struct {
	Url string `json:"href"`
}

func (n NextHref) String() string {
	return n.Url
}

type Order struct {
	Buy       bool    `json:"buy"`
	Issued    CCPTime `json:"issued"`
	Price     float64 `json:"price"`
	Volume    uint64  `json:"volume"`
	Duration  int     `json:"duration"`
	Id        uint64  `json:"id"`
	StationID uint64  `json:"stationID"`
	ItemID    uint64  `json:"type"`
}

func (o *Order) ExpiryDate() string {
	return o.Issued.AddDate(0, 0, o.Duration).Format(DB_DATETIME)
}
