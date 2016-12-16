package concatenator

import (
	"fmt"
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

func (this *PriceStats) addAvg(price float64, volume uint64) {
	this.Volume += volume
	this.Mean += (price - this.Mean) * float64(volume) / float64(this.Volume)
}

func (this *PriceStats) cmpMin(price float64) {
	if this.Min > price {
		this.Min = price
	}
}
func (this *PriceStats) cmpMax(price float64) {
	if this.Max < price {
		this.Max = price
	}
}

func (this *PriceStats) Add(price float64, volume uint64) {
	this.addAvg(price, volume)
	this.cmpMax(price)
	this.cmpMin(price)
}

func (this *PriceStats) Merge(other *PriceStats) {
	if this.Min == 0 {
		this.Min = other.Min
		this.Mean = other.Mean
		this.Max = other.Max
		this.Volume = other.Volume
		return
	}
	this.addAvg(other.Mean, other.Volume)
	this.cmpMax(other.Max)
	this.cmpMin(other.Min)
}

func (this *PriceStats) String() string {
	return fmt.Sprintf("%.2f, %.2f, %.2f", this.Min, this.Mean, this.Max)
}
