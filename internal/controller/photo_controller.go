package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	exception2 "pbi-btpns-api/internal/exception"
	model2 "pbi-btpns-api/internal/model"
	"pbi-btpns-api/internal/service"
	"pbi-btpns-api/internal/utils"
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
	var req model2.AddPhotoRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		panic(exception2.JsonParseError{Msg: "cannot parse request body"})
	}

	// validate request body
	err = p.validate.Struct(req)
	if err != nil {
		panic(exception2.ValidationError{Msg: err.Error()})
	}

	// get param
	userId := c.Param("userId")

	// get credentials
	keys := c.Keys
	credential := keys["credentials"].(map[string]string)

	if userId != credential["userId"] {
		panic(exception2.AuthorizationError{Msg: "cannot add photo other users"})
	}

	req.UserId = userId
	req.Url = utils.GetFileNameFromUrl(req.Url)

	// logout
	response := p.photoService.AddPhoto(req)
	c.JSON(200, model2.WebResponse{
		Status:  model2.Success,
		Code:    200,
		Message: "Yay, success to add new photo",
		Data:    response,
	})
}

func (p *photoControllerImpl) UpdatePhoto(c *gin.Context) {
	// bind request body to struct
	var req model2.UpdatePhotoRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		panic(exception2.JsonParseError{Msg: "cannot parse request body"})
	}

	// validate request body
	err = p.validate.Struct(req)
	if err != nil {
		panic(exception2.ValidationError{Msg: err.Error()})
	}

	// get param
	userId := c.Param("userId")
	photoId := c.Param("photoId")

	// get credentials
	keys := c.Keys
	credential := keys["credentials"].(map[string]string)

	// check credential
	if userId != credential["userId"] {
		panic(exception2.AuthorizationError{Msg: "cannot update photo other users"})
	}

	req.Id = photoId
	req.UserId = userId
	req.Url = utils.GetFileNameFromUrl(req.Url)

	// logout
	response := p.photoService.UpdatePhoto(req)
	c.JSON(200, model2.WebResponse{
		Status:  model2.Success,
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
		panic(exception2.AuthorizationError{Msg: "cannot get photo other users"})
	}

	// logout
	p.photoService.DeletePhoto(photoId, userId)
	c.JSON(200, model2.WebResponse{
		Status:  model2.Success,
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
		panic(exception2.AuthorizationError{Msg: "cannot get photo other users"})
	}

	// logout
	response := p.photoService.GetPhotoById(photoId, userId)
	c.JSON(200, model2.WebResponse{
		Status:  model2.Success,
		Code:    200,
		Message: "Yay, success to get photo",
		Data:    response,
	})
}
