package model

type Product struct {
	Common
	NameProduct  string `json:"name_product" gorm:"varchar;not_null"`
	PriceProduct int64  `json:"price_product" gorm:"varchar;not_null"`
	CategoryId   uint   `json:"category_id"`
}
