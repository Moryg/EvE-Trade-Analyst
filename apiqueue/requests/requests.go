/**
 * Shared structures between various requests
 */
package requests

import (
	. "github.com/moryg/eve_analyst/ccptime"
	. "github.com/moryg/eve_analyst/constants"
)

const (
	DB_DATE = "test"
)

type NextHref struct {
	Url string `json:"href"`
}

func (n NextHref) String() string {
	return n.Url
}

type Order struct {
	Buy       bool    `json:"buy"`
	Issued    CCPTime `json:"issued"`
	Price     float32 `json:"price"`
	Volume    int64   `json:"volume"`
	Duration  int     `json:"duration"`
	Id        int     `json:"id"`
	StationID int     `json:"stationID"`
	ItemID    int     `json:"type"`
}

func (o *Order) ExpiryDate() string {
	return o.Issued.AddDate(0, 0, o.Duration).Format(DB_DATETIME)
}
