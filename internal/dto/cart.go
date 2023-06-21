package dto

import "gorm.io/gorm"

type CartAndTotalResponse struct {
	CartResponse []*CartResponse `json:"cart"`
	TotalCart    int64           `json:"total_cart"`
}

type CartResponse struct {
	ID       uint            `json:"id"`
	Username string          `json:"username,omitempty"`
	Product  ProductResponse `json:"cart_product"`
	Quantity int64           `json:"quantity_product"`
	Price    int64           `json:"price_product"`
}

type Cart struct {
	UserId    uint
	ProductId uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
	Price     int64
}

type CartDeleteResponse struct {
	ID        uint            `json:"id"`
	ProductId uint            `json:"product_id"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}
