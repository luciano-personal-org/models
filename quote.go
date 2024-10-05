package models

// Quote struct for quote.
type Quote struct {
	ID              string
	StockName       string
	TimeStamp       string
	Date            string
	Current         float64
	Open            float64
	High            float64
	Low             float64
	Close           float64
	Volume          uint64
	VolumeVariation float64
}
