package dto

const (
	BELUM_DIBAYAR = "belum dibayar"
	SUDAH_DIBAYAR = "sudah dibayar"
)

type TransactionResponse struct {
	ID                uint                     `json:"id"`
	Product           []*ProductDetailResponse `json:"cart_product,omitempty"`
	Price             int64                    `json:"total_transaction"`
	StatusTransaction string                   `json:"status_transaction"`
}

type TransactionReq struct {
	UserId uint
	// Cart   []Cart
	CartId []uint `json:"cart_id" validate:"required"`
	Total  int64
}
