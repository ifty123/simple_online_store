package dto

import "gorm.io/gorm"

type CartResponse struct {
	ID       uint            `json:"id"`
	Username string          `json:"username,omitempty"`
	Product  ProductResponse `json:"cart_product"`
	Quantity int64           `json:"quantity_product"`
	Price    int64           `json:"price_product"`
}

type Cart struct {
	UserId    uint `json:"user_id"`
	ProductId uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
	Price     int64
}

type CartDeleteResponse struct {
	ID        uint            `json:"id"`
	ProductId uint            `json:"product_id"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}
