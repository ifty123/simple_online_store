package dto

type RegisterUserReq struct {
	Username string `json:"username" validate:"required" example:"example"`
	Email    string `json:"email" validate:"required" example:"example@gmail.com"`
	Password string `json:"password" validate:"required" example:"P@sswo4d"`
}

//example : digunakan untuk swagger
type UserResponse struct {
	ID       uint   `json:"id" example:"1"`
	Username string `json:"username" example:"example"`
	Email    string `json:"email" example:"example@gmail.com"`
}

type UserWithJWTResponse struct {
	UserResponse
	JWT string `json:"jwt"`
}
