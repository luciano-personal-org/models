package models

// Quote struct for quote.
type Signal struct {
	ID          string          `json:"ID"`              // Unique identifier for the signal
	Candle      Candle          `json:"Candle"`          // The candle that generated the signal
	Direction   SignalDirection `json:"SignalDirection"` // BUY or SELL
	PriceTarget float64         `json:"PriceTarget"`     // The price target for the signal
	StopLoss    float64         `json:"StopLoss"`        // The price for the loss of the signal
	Strength    float64         `json:"Strength"`        // 0 to 1
}

// SignalDirection is an enum-like type for signal direction.
type SignalDirection int

const (
	BUY SignalDirection = iota
	SELL
)

// String returns the string representation of the SignalDirection.
func (d SignalDirection) String() string {
	return [...]string{"BUY", "SELL"}[d]
}
