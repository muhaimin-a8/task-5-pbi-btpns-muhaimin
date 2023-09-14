package controller

import (
	"github.com/go-playground/validator/v10"
	"pbi-btpns-api/internal/service"
)

type Controllers interface {
	NewUserController() UserController
	NewAuthController() AuthController
	NewPhotoController() PhotoController
	NewUploadController() UploadController
	NewApiKeyController() ApiKeyController
}

type controllerImpl struct {
	validate      *validator.Validate
	userService   service.UserService
	authService   service.AuthService
	photoService  service.PhotoService
	apiKeyService service.ApiKeyService
}

func (c *controllerImpl) NewApiKeyController() ApiKeyController {
	return &apiKeyControllerImpl{apiKeyService: c.apiKeyService}
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
	userService service.UserService,
	authService service.AuthService,
	photoService service.PhotoService,
	apiKeyService service.ApiKeyService) Controllers {
	return &controllerImpl{
		validate:      validate,
		userService:   userService,
		authService:   authService,
		photoService:  photoService,
		apiKeyService: apiKeyService,
	}
}
