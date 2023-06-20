package model

type TransactionDetail struct {
	Common
	TransactionId uint `json:"transaction_id"`
	ProductId     uint `json:"product_id"`
	Product       Product
	Quantity      int   `json:"quantity"`
	PriceTotal    int64 `json:"price_total"`
}
