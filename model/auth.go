package model

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required,min=1"`
}

type UpdateTokenRequest struct {
	RefreshToken string `json:"access_token" validate:"required,min=1"`
}

type UpdateTokenResponse struct {
	AccessToken string `json:"access_token"`
}
