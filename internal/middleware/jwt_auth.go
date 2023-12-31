package middleware

import (
	"github.com/gin-gonic/gin"
	"pbi-btpns-api/internal/model"
	"pbi-btpns-api/internal/service"
	"strings"
)

type JwtAuthMiddleware interface {
	Init(c *gin.Context)
}

type jwtAuthMiddleware struct {
	tokenManager service.JwtTokenManager
}

func (j *jwtAuthMiddleware) Init(c *gin.Context) {
	token := getTokenFromHeader(c.GetHeader("Authorization"))
	if token == "" {
		c.AbortWithStatusJSON(401, model.WebResponse{
			Status:  model.Fail,
			Code:    401,
			Message: "invalid jwt token",
			Data:    nil,
		})

		return
	}
	userId, err := j.tokenManager.ParseAccessToken(token)
	if err != nil {
		c.AbortWithStatusJSON(401, model.WebResponse{
			Status:  model.Fail,
			Code:    401,
			Message: "invalid jwt token",
			Data:    nil,
		})

		return
	}

	c.Set("credentials", map[string]string{
		"userId": *userId,
	})
}

func getTokenFromHeader(tokenHeader string) string {
	token, _ := strings.CutPrefix(tokenHeader, "Bearer ")
	return token
}

func NewJwtAuthMiddleware(tokenManager service.JwtTokenManager) JwtAuthMiddleware {
	return &jwtAuthMiddleware{
		tokenManager: tokenManager,
	}
}
