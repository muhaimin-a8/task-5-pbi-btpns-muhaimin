package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	exception2 "pbi-btpns-api/internal/exception"
	model2 "pbi-btpns-api/internal/model"
	"pbi-btpns-api/internal/service"
)

type UserController interface {
	RegisterUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userControllerImpl struct {
	validate    *validator.Validate
	userService service.UserService
}

func (u *userControllerImpl) RegisterUser(c *gin.Context) {
	var registerRequestModel model2.UserRegisterRequest
	err := c.ShouldBindJSON(&registerRequestModel)
	if err != nil {
		panic(exception2.JsonParseError{Msg: "cannot parse request body"})
	}

	err = u.validate.Struct(registerRequestModel)
	if err != nil {
		panic(exception2.ValidationError{Msg: err.Error()})
	}
	response := u.userService.RegisterUser(registerRequestModel)
	c.JSON(201, model2.WebResponse{
		Status:  model2.Success,
		Code:    201,
		Message: "success to register new user",
		Data:    response,
	})
}

func (u *userControllerImpl) UpdateUser(c *gin.Context) {
	// bind request body
	var req model2.UserUpdateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		panic(exception2.JsonParseError{Msg: "cannot parse request body"})
	}

	// validate request body
	err = u.validate.Struct(req)
	if err != nil {
		panic(exception2.ValidationError{Msg: err.Error()})
	}

	// get param
	userId := c.Param("userId")

	// get credentials
	keys := c.Keys
	credential := keys["credentials"].(map[string]string)

	if userId != credential["userId"] {
		panic(exception2.AuthorizationError{Msg: "cannot update other users"})
	}

	req.Id = userId

	response := u.userService.UpdateUser(req)
	c.JSON(201, model2.WebResponse{
		Status:  model2.Success,
		Code:    201,
		Message: "success to update user",
		Data:    response,
	})
}

func (u *userControllerImpl) DeleteUser(c *gin.Context) {
	userId := c.Param("userId")

	// get credentials
	keys := c.Keys
	credential := keys["credentials"].(map[string]string)

	if userId != credential["userId"] {
		panic(exception2.AuthorizationError{Msg: "cannot delete other users"})
	}

	u.userService.DeleteUserById(userId)
	c.JSON(201, model2.WebResponse{
		Status:  model2.Success,
		Code:    201,
		Message: "success to delete user",
		Data:    nil,
	})
}
