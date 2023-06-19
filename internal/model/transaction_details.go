package model

type TransactionDetail struct {
	Common
	TransactionId uint `json:"transaction_id"`
	ProductId     uint `json:"product_id"`
}
