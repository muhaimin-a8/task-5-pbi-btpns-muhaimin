package model

type AddPhotoRequest struct {
	ImgId   string `json:"img_id" validate:"required,min=1,max=50"`
	Title   string `json:"title" validate:"required,min=1,max=20"`
	Caption string `json:"caption" validate:"required,min=1,max=50"`
}

type AddPhotoResponse struct {
	Id      string
	ImgId   string
	Title   string
	Caption string
	Url     string
}

type GetPhotoResponse struct {
	Id      string
	Title   string
	Caption string
	Url     string
}

type UpdatePhotoRequest struct {
	ImgId   string `json:"img_id" validate:"required,min=1,max=50"`
	Title   string `json:"title" validate:"required,min=1,max=20"`
	Caption string `json:"caption" validate:"required,min=1,max=50"`
}

type UpdatePhotoResponse struct {
	Id      string
	Title   string
	Caption string
	Url     string
}