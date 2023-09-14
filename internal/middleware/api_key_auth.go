package middleware

import (
	"github.com/gin-gonic/gin"
	"pbi-btpns-api/internal/model"
)

func (a *apiKeyMiddleware) Init(c *gin.Context) {
	apiKey := c.GetHeader("X-API-KEY")
	if apiKey != "SECRET_API_KEY" {
		c.AbortWithStatusJSON(401, model.WebResponse{
			Status:  model.Fail,
			Code:    401,
			Message: "invalid API KEY",
			Data:    nil,
		})
	}

	c.Next()
}

type apiKeyMiddleware struct {
}
