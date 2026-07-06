package dto

type AuthRegisterReq struct {
	Email    string `json:"email" binding:"email,required"`
	FullName string `json:"fullName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthCreateUserForGoogleReq struct {
	Email    string
	FullName string
	Avatar   string
	GoogleId string
}

type AuthLoginReq struct {
	Email    string `json:"email" binding:"email,required" example:"long5@gmail.com"`
	Password string `json:"password" binding:"required" example:"12345"`
}

type AuthLoginReturn struct {
	AccessToken  string
	RefreshToken string
}

type AuthRefreshTokenReturn struct {
	AccessToken  string
	RefreshToken string
}

type AuthGoogleLoginReturn struct {
	State string
	Url   string
}

type AuthGoogleCallbackInput struct {
	StateGoogle string
	StateCookie string
	Code        string
}
