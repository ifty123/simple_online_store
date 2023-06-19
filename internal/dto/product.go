package dto

type ProductResponse struct {
	ID          uint   `json:"id"`
	NameProduct string `json:"name_product"`
	Category    string `json:"category"`
	Price       int64  `json:"price"`
}
