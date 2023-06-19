package model

type Transaction struct {
	Common
	UserId            uint   `json:"user_id"`
	TotalTransaction  int    `json:"total_transaction"`
	StatusTransaction string `json:"status_transaction"`
}
