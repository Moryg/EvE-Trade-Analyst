package shared

import (
// "bytes"
// "fmt"
)

type CompItem struct {
	ID uint64 `db:"itemId" json:"id"`
	// StationID uint64  `db:"stationId" json:"station_id"`
	Name      string  `db:"name" json:"name"`
	MinSell   float64 `db:"sMin" json:"sell_min"`
	MeanSell  float64 `db:"sMean" json:"sell_mean"`
	MinBuy    float64 `db:"bMin" json:"buy_min"`
	MeanBuy   float64 `db:"bMean" json:"buy_mean"`
	MinRatio  float64 `db:"rMin" json:"ratio_min"`
	MeanRatio float64 `db:"rMean" json:"ratio_mean"`
}

type CompHolder struct {
	Items []CompItem `json:"items"`
}

// func (holder *CompHolder) MarshalJSON() ([]byte, error) {
// 	buffer := bytes.NewBufferString("{")
// 	buffer.WriteString("\"columns\":[\"itemId\",\"minSell\",\"minBuy\",\"minRatio\",\"meanSell\",\"meanBuy\",\"meanRatio\"],\"items\":[")
// 	sep := ""
// 	for _, item := range holder.Items {
// 		buffer.WriteString(fmt.Sprintf("%s[%d,%.2f,%.2f,%.2f,%.2f,%.2f,%.2f]", sep, item.ID, item.MinSell, item.MinBuy, item.MinRatio, item.MeanSell, item.MeanBuy, item.MeanRatio))
// 		sep = ","
// 	}
// 	buffer.WriteString("]}")
// 	return buffer.Bytes(), nil
// }
