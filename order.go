package models

// Orders struct for the Orders.
type Orders struct {
	ID       string      `json:"ID"`          // Unique identifier for the signal
	First    FirstOrder  `json:"FirstOrder"`  // Details of the First Signal
	Second   SecondOrder `json:"SecondOrder"` // Details of the Second Signal
	Strength float64     `json:"Strength"`    // 0.0 to 1.0
	Candle   Candle      `json:"Candle"`      // The candle that generated the signal
	Volume   float64     `json:"Volume"`      // The volume of the signal
}

// FirstOrder struct for the First Order.
type FirstOrder struct {
	ID        string    `json:"ID"`        // Unique identifier for the signal
	OrderSide OrderSide `json:"OrderSide"` // The side of order
	OrderType OrderType `json:"OrderType"` // The type of order
	Price     float64   `json:"Price"`     // The price of the order
}

// SecondOrder struct for the Second Order.
type SecondOrder struct {
	ID        string    `json:"ID"`        // Unique identifier for the signal
	StopGain  float64   `json:"StopGain"`  // The price target for the signal
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
