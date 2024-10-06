package models

type Book struct {
	Id           string  `json:"ID"`
	StockName    string  `json:"StockName"`
	BookType     byte    `json:"BookType"`
	Position     uint    `json:"Position"`
	Direction    byte    `json:"Direction"`
	Price        float64 `json:"Price"`
	Date         string  `json:"Date"`
	Time         string  `json:"Time"`
	OrderId      string  `json:"OrderId"`
	Quantity     uint    `json:"Quantity"`
	OfferType    byte    `json:"OfferType"`
	QuantityMean uint64  `json:"QuantityMean"`
}
