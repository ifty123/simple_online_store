package dto

type CartResponse struct {
	ID       uint              `json:"id"`
	Username string            `json:"username"`
	Product  []ProductResponse `json:"cart_product"`
	Quantity int64             `json:"quantity_product"`
	Price    int64             `json:"price_product"`
}
