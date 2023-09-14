package controller

import (
	"github.com/go-playground/validator/v10"
	service2 "pbi-btpns-api/internal/service"
)

type Controllers interface {
	NewUserController() UserController
	NewAuthController() AuthController
	NewPhotoController() PhotoController
	NewUploadController() UploadController
}

type controllerImpl struct {
	validate     *validator.Validate
	userService  service2.UserService
	authService  service2.AuthService
	photoService service2.PhotoService
}

func (c *controllerImpl) NewUploadController() UploadController {
	return &uploadControllerImpl{}
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
	return &photoControllerImpl{
		validate:     c.validate,
		photoService: c.photoService,
	}
}

func NewController(validate *validator.Validate,
	userService service2.UserService,
	authService service2.AuthService,
	photoService service2.PhotoService) Controllers {
	return &controllerImpl{
		validate:     validate,
		userService:  userService,
		authService:  authService,
		photoService: photoService,
	}
}
