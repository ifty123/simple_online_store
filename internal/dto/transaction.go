package dto

type TransactionResponse struct {
	ID                uint              `json:"id"`
	Username          string            `json:"username"`
	Product           []ProductResponse `json:"cart_product"`
	Price             int64             `json:"total_transaction"`
	StatusTransaction string            `json:"status_transaction"`
}
