package middleware

import (
	"pbi-btpns-api/internal/service"
)

type Middlewares interface {
	NewJwtAuth() JwtAuthMiddleware
	NewApiKey() *apiKeyMiddleware
}

type middlewaresImpl struct {
	tokenManager service.JwtTokenManager
}

func (m *middlewaresImpl) NewJwtAuth() JwtAuthMiddleware {
	return &jwtAuthMiddleware{tokenManager: m.tokenManager}
}

func (m *middlewaresImpl) NewApiKey() *apiKeyMiddleware {
	return &apiKeyMiddleware{}
}

func NewMiddlewares(tokenManager service.JwtTokenManager) Middlewares {
	return &middlewaresImpl{
		tokenManager: tokenManager,
	}
}
