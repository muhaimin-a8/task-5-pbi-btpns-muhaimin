package controller

import (
	"github.com/go-playground/validator/v10"
	"pbi-btpns-api/service"
)

type Controllers interface {
	NewUserController() UserController
	NewAuthController() AuthController
	NewPhotoController() PhotoController
}

type controllerImpl struct {
	validate     *validator.Validate
	userService  service.UserService
	authService  service.AuthService
	photoService service.PhotoService
}

func (c *controllerImpl) NewUserController() UserController {
	return &userControllerImpl{
		validate:    c.validate,
		userService: c.userService,
	}
}

func (c *controllerImpl) NewAuthController() AuthController {
	return &authControllerImpl{
		validate:    c.validate,
		authService: c.authService,
	}
}

func (c *controllerImpl) NewPhotoController() PhotoController {
	//TODO implement me
	panic("implement me")
}

func NewController(validate *validator.Validate,
	userService service.UserService,
	authService service.AuthService,
	photoService service.PhotoService) Controllers {
	return &controllerImpl{
		validate:     validate,
		userService:  userService,
		authService:  authService,
		photoService: photoService,
	}
}
