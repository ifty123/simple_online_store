package model

type User struct {
	Common
	Username string `json:"username" gorm:"varchar;not_null"`
	Email    string `json:"email" gorm:"varchar;not_null;unique"`
	Password string `json:"password" gorm:"varchar;not_null"`
}
