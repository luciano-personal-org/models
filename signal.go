package models

// Quote struct for the Signal.
type Signal struct {
	ID        string          `json:"ID"`              // Unique identifier for the signal
	Candle    Candle          `json:"Candle"`          // The candle that generated the signal
	Direction SignalDirection `json:"SignalDirection"` // BUY or SELL
	Price     float64         `json:"Price"`           // The price target for the signal
	StopLoss  float64         `json:"StopLoss"`        // The price for the loss of the signal
	Strength  float64         `json:"Strength"`        // 0.0 to 1.0
	OrderType OrderType       `json:"OrderType"`       // The type of order
}

// SignalDirection is an enum-like type for signal direction.
type SignalDirection int

const (
	SIGNAL_BUY SignalDirection = iota
	SIGNAL_SELL
)

// String returns the string representation of the SignalDirection.
func (d SignalDirection) String() string {
	return [...]string{"SIGNAL_BUY", "SIGNAL_SELL"}[d]
}
