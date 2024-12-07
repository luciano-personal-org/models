package models

// Quote struct for quote.
type Signal struct {
	ID          string          `json:"ID"`              // Unique identifier for the signal
	Candle      Candle          `json:"Candle"`          // The candle that generated the signal
	Direction   SignalDirection `json:"SignalDirection"` // BUY or SELL
	PriceOrigin float64         `json:"PriceOrigin"`     // The price that originated the signal
	PriceTarget float64         `json:"PriceTarget"`     // The price target for the signal
	StopLoss    float64         `json:"StopLoss"`        // The price for the loss of the signal
	Strength    float64         `json:"Strength"`        // 0.0 to 1.0
	OrderType   OrderType       `json:"OrderType"`       // The type of order
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

type OrderType int

const (
	LIMIT OrderType = iota
	MARKET
	STOP_LOSS
	STOP_LOSS_LIMIT
	TAKE_PROFIT
	TAKE_PROFIT_LIMIT
	LIMIT_MAKER
)

// String returns the string representation of the SignalDirection.
func (d OrderType) String() string {
	return [...]string{
		"LIMIT",
		"MARKET",
		"STOP_LOSS",
		"STOP_LOSS_LIMIT",
		"TAKE_PROFIT",
		"TAKE_PROFIT_LIMIT",
		"LIMIT_MAKER",
	}[d]
}
