package middleware

import (
	"github.com/gin-gonic/gin"
	"pbi-btpns-api/internal/model"
	"pbi-btpns-api/internal/service"
)

func (a *apiKeyMiddleware) Init(c *gin.Context) {
	apiKey := c.GetHeader("X-API-KEY")
	isExist := a.apiKeyService.IsExist(apiKey)
	if !isExist {
		c.AbortWithStatusJSON(401, model.WebResponse{
			Status:  model.Fail,
			Code:    401,
			Message: "Invalid API KEY",
			Data:    nil,
		})
	}
}

type apiKeyMiddleware struct {
	apiKeyService service.ApiKeyService
}
