package models

type Trade struct {
	ID          string      `json:"ID"`          // Unique identifier for the signal
	FristOrder  FirstOrder  `json:"FirstOrder"`  // Details of the First Signal
	SecondOrder SecondOrder `json:"SecondOrder"` // Details of the Second Signal
	Quantity    float64     `json:"Quantity"`    // The quantity of the order

}
