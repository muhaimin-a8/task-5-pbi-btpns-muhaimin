package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"pbi-btpns-api/exception"
	"pbi-btpns-api/model"
	"pbi-btpns-api/service"
)

type AuthController interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	UpdateAccessToken(c *gin.Context)
}

type authControllerImpl struct {
	validate    *validator.Validate
	authService service.AuthService
}

func (a *authControllerImpl) Login(c *gin.Context) {
	// bind request body to struct
	var loginReq model.LoginRequest
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		panic(exception.JsonParseError{Msg: "cannot parse request body"})
	}

	// validate request body
	err = a.validate.Struct(loginReq)
	if err != nil {
		panic(exception.ValidationError{Msg: err.Error()})
	}

	// login
	response := a.authService.Login(loginReq)
	c.JSON(200, model.WebResponse{
		Status:  model.Success,
		Code:    200,
		Message: "Yay, success to login",
		Data:    response,
	})
}

func (a *authControllerImpl) Logout(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *authControllerImpl) UpdateAccessToken(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}