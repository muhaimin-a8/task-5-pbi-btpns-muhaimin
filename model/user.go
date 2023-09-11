package model

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required,min=1,max=25"`
	Email    string `json:"email" validate:"required,email,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}
type UserRegisterResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserUpdateRequest struct {
	Id       string
	Username string `json:"username" validate:"required,min=1,max=25"`
	Email    string `json:"email" validate:"required,email,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type UserUpdateResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
