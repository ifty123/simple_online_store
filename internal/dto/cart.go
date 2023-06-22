package dto

import (
	"time"

	"gorm.io/gorm"
)

// example: digunakan untuk swagger documentation
type CartAndTotalResponse struct {
	CartResponse []*CartResponse `json:"cart"`
	TotalCart    int64           `json:"total_cart" example:"80000"`
}

// example: digunakan untuk swagger documentation
type CartResponse struct {
	ID       uint            `json:"id" example:"1"`
	Product  ProductResponse `json:"cart_product"`
	Quantity int64           `json:"quantity_product" example:"1"`
	Price    int64           `json:"price_product" example:"80000"`
}

type Cart struct {
	UserId    uint  `json:"-"`
	ProductId uint  `json:"product_id" example:"1"`
	Quantity  int   `json:"quantity" example:"1"`
	Price     int64 `json:"-"`
}

type CartDeleteResponse struct {
	ID        uint      `json:"id" example:"1"`
	ProductId uint      `json:"product_id" example:"1"`
	Deleted   time.Time `json:"deleted_at"`
	// @Ignore
	DeletedAt *gorm.DeletedAt `json:"-"`
}
