package shared

import (
	"bytes"
	"fmt"
)

type CompItem struct {
	ID        uint64  `db:"itemId"`
	StationID uint64  `db:"stationId"`
	Name      string  `db:"name"`
	MinSell   float64 `db:"sMin"`
	MeanSell  float64 `db:"sMean"`
	MinBuy    float64 `db:"bMin"`
	MeanBuy   float64 `db:"bMean"`
	MinRatio  float64 `db:"rMin"`
	MeanRatio float64 `db:"rMean"`
}

type CompHolder struct {
	Items []CompItem
}

func (holder *CompHolder) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	buffer.WriteString("\"columns\":[\"itemId\",\"minSell\",\"minBuy\",\"minRatio\",\"meanSell\",\"meanBuy\",\"meanRatio\"],\"items\":[")
	sep := ""
	for _, item := range holder.Items {
		buffer.WriteString(fmt.Sprintf("%s[%d,%.2f,%.2f,%.2f,%.2f,%.2f,%.2f]", sep, item.ID, item.MinSell, item.MinBuy, item.MinRatio, item.MeanSell, item.MeanBuy, item.MeanRatio))
		sep = ","
	}
	buffer.WriteString("]}")
	return buffer.Bytes(), nil
}
