package models

type Book struct {
	Id           string  `json:"ID"`
	StockName    string  `json:"StockName"`
	BookType     byte    `json:"BookType"`
	Position     uint    `json:"Position"`
	Direction    byte    `json:"Direction"`
	Price        float64 `json:"Price"`
	DateTime     string  `json:"DateTime"`
	OrderId      string  `json:"OrderId"`
	Quantity     uint    `json:"Quantity"`
	OfferType    byte    `json:"OfferType"`
	QuantityMean uint    `json:"QuantityMean"`
}
