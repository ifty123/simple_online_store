package dto

type ProductResponse struct {
	ID          uint   `json:"id"`
	NameProduct string `json:"name_product"`
	Category    string `json:"category,omitempty"`
	Price       int64  `json:"price"`
}

type ProductDetailResponse struct {
	ID          uint   `json:"id"`
	NameProduct string `json:"name_product"`
	Price       int64  `json:"price,omitempty"`
	Quantity    int    `json:"quantity"`
	PriceTotal  int64  `json:"price_total,omitempty"`
}
