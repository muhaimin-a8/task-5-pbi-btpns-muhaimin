package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	exception2 "pbi-btpns-api/internal/exception"
	model2 "pbi-btpns-api/internal/model"
	"pbi-btpns-api/internal/service"
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
	var loginReq model2.LoginRequest
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		panic(exception2.JsonParseError{Msg: "cannot parse request body"})
	}

	// validate request body
	err = a.validate.Struct(loginReq)
	if err != nil {
		panic(exception2.ValidationError{Msg: err.Error()})
	}

	// login
	response := a.authService.Login(loginReq)
	c.JSON(200, model2.WebResponse{
		Status:  model2.Success,
		Code:    200,
		Message: "Yay, success to login",
		Data:    response,
	})
}

func (a *authControllerImpl) Logout(c *gin.Context) {
	// bind request body to struct
	var req model2.LogoutRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		panic(exception2.JsonParseError{Msg: "cannot parse request body"})
	}

	// validate request body
	err = a.validate.Struct(req)
	if err != nil {
		panic(exception2.ValidationError{Msg: err.Error()})
	}

	// logout
	a.authService.Logout(req)
	c.JSON(200, model2.WebResponse{
		Status:  model2.Success,
		Code:    200,
		Message: "Yay, success to logout",
		Data:    nil,
	})
}

func (a *authControllerImpl) UpdateAccessToken(c *gin.Context) {
	// bind request body to struct
	var req model2.UpdateTokenRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		panic(exception2.JsonParseError{Msg: "cannot parse request body"})
	}

	// validate request body
	err = a.validate.Struct(req)
	if err != nil {
		panic(exception2.ValidationError{Msg: err.Error()})
	}

	// logout
	response := a.authService.UpdateToken(req)
	c.JSON(200, model2.WebResponse{
		Status:  model2.Success,
		Code:    200,
		Message: "Yay, success to update token",
		Data:    response,
	})
}
