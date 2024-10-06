package models

// Quote struct for quote.
type Candle struct {
	ID              string  `json:"ID"`
	StockName       string  `json:"StockName"`
	TimeStamp       string  `json:"TimeStamp"`
	Date            string  `json:"Date"`
	Current         float64 `json:"Current"`
	Open            float64 `json:"Open"`
	High            float64 `json:"High"`
	Low             float64 `json:"Low"`
	Close           float64 `json:"Close"`
	Volume          uint64  `json:"Volume"`
	VolumeVariation float64 `json:"VolumeVariation"`
}
