package dto

// example: digunakan untuk documentation swagger
type ProductResponse struct {
	ID          uint   `json:"id" example:"1"`
	NameProduct string `json:"name_product" example:"Kemeja Denim"`
	Category    string `json:"category,omitempty" example:"Kemeja"`
	Price       int64  `json:"price" example:"80000"`
}

type ProductDetailResponse struct {
	ID          uint   `json:"id" example:"1"`
	NameProduct string `json:"name_product" example:"Kemeja Denim"`
	Price       int64  `json:"price,omitempty" example:"80000"`
	Quantity    int    `json:"quantity" example:"4"`
	PriceTotal  int64  `json:"price_total,omitempty" example:"320000"`
}
