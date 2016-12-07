package concatenator

import ()

type Region struct {
	Prices map[uint64]map[uint64]*PriceStats
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
