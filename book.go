package models

type Book struct {
	Id        string  `json:"ID"`
	StockName string  `json:"StockName"`
	BookType  byte    `json:"BookType"`
	Position  uint    `json:"Position"`
	Direction byte    `json:"Direction"`
	Price     float64 `json:"Price"`
	Quantity  uint    `json:"Quantity"`
	DateTime  string  `json:"DateTime"`
	OrderId   string  `json:"OrderId"`
}
