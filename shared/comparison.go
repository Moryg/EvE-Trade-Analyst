package shared

import (
	"bytes"
	"fmt"
)

type CompItem struct {
	ID        int     `db:"id"`
	GroupID   int     `db:"groupID"`
	Volume    float64 `db:"volume"`
	Name      string  `db:"name"`
	MinSell   float64 `db:"sellMin"`
	MinBuy    float64 `db:"buyMin"`
	MinRatio  float64 `db:"minRatio"`
	MeanSell  float64 `db:"sellMean"`
	MeanBuy   float64 `db:"buyMean"`
	MeanRatio float64 `db:"meanRatio"`
}

type CompHolder struct {
	Items []CompItem
}

func (holder *CompHolder) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	buffer.WriteString("\"columns\":[\"itemId\",\"minSell\",\"minBuy\",\"minRatio\",\"meanSell\",\"meanBuy\",\"meanRatio\"],\"items\":[")
	sep := ""
	for _, item := range holder.Items {
		buffer.WriteString(sep + fmt.Sprintf("[%d,%.2f,%.2f,%.2f,%.2f,%.2f,%.2f]", item.ID, item.MinSell, item.MinBuy, item.MinRatio, item.MeanSell, item.MeanBuy, item.MeanRatio))
		sep = ","
	}
	buffer.WriteString("]}")
	return buffer.Bytes(), nil
}
