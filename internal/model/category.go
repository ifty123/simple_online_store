package model

type Category struct {
	Common
	NameCategory string `json:"name_category" gorm:"varchar;not_null"`
}
