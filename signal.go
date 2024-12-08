package models

// Quote struct for the Signal.
type Signals struct {
	ID     string `json:"ID"`     // Unique identifier for the signal
	First  Signal `json:"First"`  // The candle that generated the signal
	Second Signal `json:"Second"` // The candle that generated the signal
}

type Signal struct {
	ID        string          `json:"ID"`              // Unique identifier for the signal
	Candle    Candle          `json:"Candle"`          // The candle that generated the signal
	Direction SignalDirection `json:"SignalDirection"` // BUY or SELL
	StopGain  float64         `json:"StopGain"`        // The price target for the signal
	StopLoss  float64         `json:"StopLoss"`        // The price for the loss of the signal
	Strength  float64         `json:"Strength"`        // 0.0 to 1.0
	OrderType OrderType       `json:"OrderType"`       // The type of order
}

// SignalDirection is an enum-like type for signal direction.
type SignalDirection int

const (
	BULLISH SignalDirection = iota
	BEARISH
)

// String returns the string representation of the SignalDirection.
func (d SignalDirection) String() string {
	return [...]string{"BULLISH", "BEARISH"}[d]
}
