package models

// Orders struct for the Orders.
type Orders struct {
	ID       string      `json:"ID"`          // Unique identifier for the signal
	First    FirstOrder  `json:"FirstOrder"`  // Details of the First Signal
	Second   SecondOrder `json:"SecondOrder"` // Details of the Second Signal
	Strength float64     `json:"Strength"`    // 0.0 to 1.0
	Candle   Candle      `json:"Candle"`      // The candle that generated the signal
}

// FirstOrder struct for the First Order.
type FirstOrder struct {
	ID        string    `json:"ID"`        // Unique identifier for the signal
	OrderSide OrderSide `json:"OrderSide"` // The side of order
	OrderType OrderType `json:"OrderType"` // The type of order
}

// SecondOrder struct for the Second Order.
type SecondOrder struct {
	ID        string    `json:"ID"`        // Unique identifier for the signal
	StopGain  float64   `json:"StopGain"`  // The price target for the signal
	StopLoss  float64   `json:"StopLoss"`  // The price for the loss of the signal
	Exit10    float64   `json:"Exit10"`    // The price for the 10% exit of the signal
	Exit20    float64   `json:"Exit20"`    // The price for the 20% exit of the signal
	Exit30    float64   `json:"Exit30"`    // The price for the 30% exit of the signal
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

type TypeInForce int

const (
	GTC TypeInForce = iota
)

func (d TypeInForce) String() string {
	return [...]string{"GTC"}[d]
}

type NewOrderRespType int

const (
	ACK NewOrderRespType = iota
	FULL
	RESULT
)

func (d NewOrderRespType) String() string {
	return [...]string{"ACK", "FULL", "RESULT"}[d]
}
