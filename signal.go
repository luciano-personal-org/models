package models

// Quote struct for quote.
type Signal struct {
	Candle      Candle          // The candle that generated the signal
	Direction   SignalDirection // BUY or SELL
	PriceTarget float64         // The price target for the signal
	StopLoss    float64         // The price for the loss of the signal
	Strength    float64         // 0 to 1
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
