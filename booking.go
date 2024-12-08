package models

// Quote struct for the Order.
type Booking struct {
	ID              string          `json:"ID"`              // Unique identifier for the signal
	Signals         Orders          `json:"Orders"`          // The candle that generated the signal
	Quantity        float64         `json:"Quantity"`        // The quantity of the order
	MarketDirection MarketDirection `json:"MarketDirection"` // The type of order
}

type MarketDirection int

const (
	BULLISH MarketDirection = iota
	BEARISH
)

// String returns the string representation of the SignalDirection.
func (d MarketDirection) String() string {
	return [...]string{"BULLISH", "BEARISH"}[d]
}
