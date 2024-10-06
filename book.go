package models

type Book struct {
	Id        string  `json:"ID"`
	StockName string  `json:"StockName"`
	Position  uint    `json:"Position"`
	Direction byte    `json:"Direction"`
	Price     float64 `json:"Price"`
	Quantity  float64 `json:"Quantity"`
	DateTime  string  `json:"DateTime"`
	OrderId   string  `json:"OrderId"`
}
