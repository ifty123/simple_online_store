package dto

const (
	BELUM_DIBAYAR = "belum dibayar"
	SUDAH_DIBAYAR = "sudah dibayar"
)

type TransactionResponse struct {
	ID                uint                     `json:"id" example:"1"`
	Product           []*ProductDetailResponse `json:"cart_product,omitempty"`
	Price             int64                    `json:"total_transaction" example:"320000"`
	StatusTransaction string                   `json:"status_transaction" example:"belum dibayar"`
}

type TransactionReq struct {
	UserId uint `json:"-"`
	// Cart   []Cart
	CartId []uint `json:"cart_id" validate:"required"`
	Total  int64  `json:"-"`
}
