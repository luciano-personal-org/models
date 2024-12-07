package models

// Quote struct for the Order.
type Order struct {
	ID        string    `json:"ID"`        // Unique identifier for the signal
	Signal    Signal    `json:"Signal"`    // The candle that generated the signal
	Side      OrderSide `json:"OrderSide"` // BUY or SELL
	Quantity  float64   `json:"Quantity"`  // The quantity of the order
	StopGain  float64   `json:"StopGain"`  // The price target for the signal
	StopLoss  float64   `json:"StopLoss"`  // The price for the loss of the signal
	Strength  float64   `json:"Strength"`  // 0.0 to 1.0
	OrderType OrderType `json:"OrderType"` // The type of order
}

// SignalDirection is an enum-like type for signal direction.
type OrderSide int

const (
	SIDE_BUY OrderSide = iota
	SIDE_SELL
)

// String returns the string representation of the SignalDirection.
func (d OrderSide) String() string {
	return [...]string{"SIDE_BUY", "SIDE_SELL"}[d]
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
