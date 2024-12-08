package models

// Orders struct for the Orders.
type Orders struct {
	ID       string      `json:"ID"`          // Unique identifier for the signal
	First    FirstOrder  `json:"FirstOrder"`  // Details of the First Signal
	Second   SecondOrder `json:"SecondOrder"` // Details of the Second Signal
	Strength float64     `json:"Strength"`    // 0.0 to 1.0
}

// FirstOrder struct for the First Order.
type FirstOrder struct {
	ID        string    `json:"ID"`        // Unique identifier for the signal
	Candle    Candle    `json:"Candle"`    // The candle that generated the signal
	OrderSide OrderSide `json:"OrderSide"` // The side of order
	OrderType OrderType `json:"OrderType"` // The type of order
}

// SecondOrder struct for the Second Order.
type SecondOrder struct {
	ID        string    `json:"ID"`        // Unique identifier for the signal
	Candle    Candle    `json:"Candle"`    // The candle that generated the signal
	StopGain  float64   `json:"StopGain"`  // The price target for the signal
	StopLoss  float64   `json:"StopLoss"`  // The price for the loss of the signal
	OrderSide OrderSide `json:"OrderSide"` // The side of order
	OrderType OrderType `json:"OrderType"` // The type of order
}

type OrderSide int

const (
	BUY OrderSide = iota
	SELL
)

func (d OrderSide) String() string {
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
