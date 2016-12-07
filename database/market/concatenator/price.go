package concatenator

import (
// "math"
)

type PriceStats struct {
	Mean, Min, Max float64
	Volume         uint64
}

func NewPrice(price float64, volume uint64) *PriceStats {
	p := new(PriceStats)
	p.Volume = volume
	p.Mean = price
	p.Min = price
	p.Max = price
	return p
}

func (p *PriceStats) Add(price float64, volume uint64) {
	p.Volume += volume
	p.Mean += (price - p.Mean) * float64(volume) / float64(p.Volume)
	if p.Max < price {
		p.Max = price
	}
	if p.Min > price {
		p.Min = price
	}
}
