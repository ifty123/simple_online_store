package dto

type RegisterUserReq struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserWithJWTResponse struct {
	UserResponse
	JWT string `json:"jwt"`
}
