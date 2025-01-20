package models

type ExecutedOrders struct {
	ID           string       `json:"ID"`           // Unique identifier for the signal
	Orders       Orders       `json:"Orders"`       // Details of the Orders
	BrokerOrders BrokerOrders `json:"BrokerOrders"` // Details of the Broker Orders
}

type BrokerOrders struct {
	FirstExecutedOrder  FirstExecutedOrder  `json:"FirstExecutedOrder"`  // Details of the First Executed Order
	SecondExecutedOrder SecondExecutedOrder `json:"SecondExecutedOrder"` // Details of the Second Executed Order
}

type FirstExecutedOrder struct {
	ID                string            `json:"ID"`            // Unique identifier for the signal
	ClientOrderId     string            `json:"ClientOrderId"` // The client order id
	TimeStamp         int64             `json:"TimeStamp"`     // The time stamp of the order
	Price             string            `json:"Price"`         // The price of the order
	OrigQty           string            `json:"OrigQty"`       // The original quantity of the order
	Status            string            `json:"Status"`        // The status of the order
	WorkingTime       int64             `json:"workingTime"`   // The time the order was placed
	Commission        string            `json:"Commission"`    // The commission of the order
	TradeId           int64             `json:"tradeId"`       // The trade id of the order
	ExecutedOrderSide ExecutedOrderSide `json:"OrderSide"`     // The side of order
}

type SecondExecutedOrder struct {
	ID                string            `json:"ID"`            // Unique identifier for the signal
	ClientOrderId     string            `json:"ClientOrderId"` // The client order id
	TimeStamp         int64             `json:"TimeStamp"`     // The time stamp of the order
	Price             string            `json:"Price"`         // The price of the order
	OrigQty           string            `json:"OrigQty"`       // The original quantity of the order
	Status            string            `json:"Status"`        // The status of the order
	WorkingTime       int64             `json:"workingTime"`   // The time the order was placed
	ExecutedOrderSide ExecutedOrderSide `json:"OrderSide"`     // The side of order
}

type ExecutedOrderSide int

const (
	EXECUTED_BUY ExecutedOrderSide = iota
	EXECUTED_SELL
)

func (d ExecutedOrderSide) String() string {
	return [...]string{"BUY", "SELL"}[d]
}
