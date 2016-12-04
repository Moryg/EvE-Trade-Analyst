package shared

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
