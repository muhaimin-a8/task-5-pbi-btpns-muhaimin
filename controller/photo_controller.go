package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"pbi-btpns-api/exception"
	"pbi-btpns-api/model"
	"pbi-btpns-api/service"
)

type PhotoController interface {
	AddPhoto(c *gin.Context)
	UpdatePhoto(c *gin.Context)
	DeletePhoto(c *gin.Context)
	GetPhoto(c *gin.Context)
}

type photoControllerImpl struct {
	validate     *validator.Validate
	photoService service.PhotoService
}

func (p *photoControllerImpl) AddPhoto(c *gin.Context) {
	// bind request body to struct
	var req model.AddPhotoRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		panic(exception.JsonParseError{Msg: "cannot parse request body"})
	}

	// validate request body
	err = p.validate.Struct(req)
	if err != nil {
		panic(exception.ValidationError{Msg: err.Error()})
	}

	// get param
	userId := c.Param("userId")

	// get credentials
	keys := c.Keys
	credential := keys["credentials"].(map[string]string)

	if userId != credential["userId"] {
		panic(exception.AuthorizationError{Msg: "cannot add photo other users"})
	}

	req.UserId = userId

	// logout
	response := p.photoService.AddPhoto(req)
	c.JSON(200, model.WebResponse{
		Status:  model.Success,
		Code:    200,
		Message: "Yay, success to add new photo",
		Data:    response,
	})
}

func (p *photoControllerImpl) UpdatePhoto(c *gin.Context) {
	// bind request body to struct
	var req model.UpdatePhotoRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		panic(exception.JsonParseError{Msg: "cannot parse request body"})
	}

	// validate request body
	err = p.validate.Struct(req)
	if err != nil {
		panic(exception.ValidationError{Msg: err.Error()})
	}

	// get param
	userId := c.Param("userId")
	photoId := c.Param("photoId")

	// get credentials
	keys := c.Keys
	credential := keys["credentials"].(map[string]string)

	// check credential
	if userId != credential["userId"] {
		panic(exception.AuthorizationError{Msg: "cannot update photo other users"})
	}

	req.Id = photoId
	req.UserId = userId

	// logout
	response := p.photoService.UpdatePhoto(req)
	c.JSON(200, model.WebResponse{
		Status:  model.Success,
		Code:    200,
		Message: "Yay, success to update new photo",
		Data:    response,
	})
}

func (p *photoControllerImpl) DeletePhoto(c *gin.Context) {
	// get param
	userId := c.Param("userId")
	photoId := c.Param("photoId")

	// get credentials
	keys := c.Keys
	credential := keys["credentials"].(map[string]string)

	if userId != credential["userId"] {
		panic(exception.AuthorizationError{Msg: "cannot get photo other users"})
	}

	// logout
	p.photoService.DeletePhoto(photoId, userId)
	c.JSON(200, model.WebResponse{
		Status:  model.Success,
		Code:    200,
		Message: "Yay, success to delete photo",
		Data:    nil,
	})
}

func (p *photoControllerImpl) GetPhoto(c *gin.Context) {
	// get param
	userId := c.Param("userId")
	photoId := c.Param("photoId")

	// get credentials
	keys := c.Keys
	credential := keys["credentials"].(map[string]string)

	if userId != credential["userId"] {
		panic(exception.AuthorizationError{Msg: "cannot get photo other users"})
	}

	// logout
	response := p.photoService.GetPhotoById(photoId, userId)
	c.JSON(200, model.WebResponse{
		Status:  model.Success,
		Code:    200,
		Message: "Yay, success to get photo",
		Data:    response,
	})
}
