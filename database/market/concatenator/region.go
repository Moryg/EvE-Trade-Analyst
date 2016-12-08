package concatenator

import (
	"errors"
	"fmt"
	"strings"
)

type Region struct {
	Prices   map[uint64]map[uint64]*PriceStats
	RegionID int
}

func NewRegion() *Region {
	this := new(Region)
	this.Prices = make(map[uint64]map[uint64]*PriceStats)
	return this
}

func (this *Region) Add(price float64, volume, stationID, itemID uint64) {
	station, ok := this.Prices[stationID]
	if !ok {
		station = make(map[uint64]*PriceStats)
		this.Prices[stationID] = station
	}

	obj, ok := station[itemID]
	if !ok {
		this.Prices[stationID][itemID] = NewPrice(price, volume)
		return
	}

	obj.Add(price, volume)
}

func (this *Region) UnmarshalJSON(b []byte) error {
	this.Prices = make(map[uint64]map[uint64]*PriceStats)
	return newParser(this).parse(b)
}

func (this *Region) Merge(other *Region) {
	for station, miniMap := range other.Prices {
		for item, otherPrice := range miniMap {
			myPrice, ok := this.Prices[station][item]
			if !ok {
				this.Add(0, 0, station, item)
				myPrice, _ = this.Prices[station][item]
			}

			myPrice.Merge(otherPrice)
		}
	}
}

func (this *Region) ConstructSQL() (string, error) {
	sql := ""
	empty := true
	if this.RegionID == 0 {
		return "", errors.New("Region ID not set")
	}

	for stationId, station := range this.Prices {
		empty = false

		for itemId, price := range station {
			// (`stationId`, `itemId`, `regionId`, `min`, `mean`, `max`, `upFlag`)
			sql += fmt.Sprintf("(%d, %d, %d, %s, 1),", stationId, itemId, this.RegionID, price.String())
		}
	}

	if empty {
		return "", errors.New("No orders in dataset")
	}

	return strings.TrimSuffix(sql, ","), nil
}
