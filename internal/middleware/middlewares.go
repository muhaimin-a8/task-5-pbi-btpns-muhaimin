package middleware

import (
	"pbi-btpns-api/internal/service"
)

type Middlewares interface {
	NewJwtAuth() JwtAuthMiddleware
	NewApiKeyAuth() *apiKeyMiddleware
}

type middlewaresImpl struct {
	tokenManager  service.JwtTokenManager
	apiKeyService service.ApiKeyService
}

func (m *middlewaresImpl) NewJwtAuth() JwtAuthMiddleware {
	return &jwtAuthMiddleware{tokenManager: m.tokenManager}
}

func (m *middlewaresImpl) NewApiKeyAuth() *apiKeyMiddleware {
	return &apiKeyMiddleware{apiKeyService: m.apiKeyService}
}

func NewMiddlewares(tokenManager service.JwtTokenManager, apiKeyService service.ApiKeyService) Middlewares {
	return &middlewaresImpl{
		tokenManager:  tokenManager,
		apiKeyService: apiKeyService,
	}
}
