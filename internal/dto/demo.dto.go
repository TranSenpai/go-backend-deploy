package dto

type DemoBody struct {
	Email     string `json:"email" binding:"required,email"`
	Passsword string `json:"password" binding:"required"`
}
