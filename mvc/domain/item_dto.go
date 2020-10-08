package domain

type Item struct {
	ItemId string  `json:"itemId"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
}
