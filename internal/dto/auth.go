package dto

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type EmailAndPasswordReq struct {
	Email    string `json:"email" validate:"required,email" example:"example@gmail.com"`
	Password string `json:"password" validate:"required" example:"P@sswo4d"`
}
