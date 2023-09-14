package controller

import (
	"github.com/gin-gonic/gin"
	"pbi-btpns-api/internal/service"
)

type ApiKeyController interface {
	RegisterNewApiKey(c *gin.Context)
}

type apiKeyControllerImpl struct {
	apiKeyService service.ApiKeyService
}

func (a *apiKeyControllerImpl) RegisterNewApiKey(c *gin.Context) {
	response := a.apiKeyService.AddNewKey()
	c.JSON(200, response)
}
