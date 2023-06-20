package model

type Cart struct {
	Common
	UserId     uint  `json:"user_id"`
	ProductId  uint  `json:"product_id"`
	Quantity   int   `json:"quantity"`
	PriceTotal int64 `json:"price_total"`
	Product    Product
}
