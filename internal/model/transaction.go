package model

type Transaction struct {
	Common
	UserId            uint   `json:"user_id"`
	TotalTransaction  int64  `json:"total_transaction"`
	StatusTransaction string `json:"status_transaction"`
	TransactionDetail []TransactionDetail
}
