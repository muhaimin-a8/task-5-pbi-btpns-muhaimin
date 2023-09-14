package service

import (
	"pbi-btpns-api/internal/model"
	"pbi-btpns-api/internal/repository"
)

type ApiKeyService interface {
	AddNewKey() *model.AddApiKeyResponse
	IsExist(key string) bool
}

type apiKeyServiceImpl struct {
	dao         repository.DAO
	idGenerator IdGenerator
}

func (a *apiKeyServiceImpl) AddNewKey() *model.AddApiKeyResponse {
	key := a.idGenerator.New(255)

	err := a.dao.NewApiKeyRepository().AddKey(key)
	if err != nil {
		panic(err)
	}

	return &model.AddApiKeyResponse{Key: key}
}

func (a *apiKeyServiceImpl) IsExist(key string) bool {
	err := a.dao.NewApiKeyRepository().VerifyKeyIsExist(key)
	if err != nil {
		return false
	}

	return true
}

func NewApiKeyService(dao repository.DAO, generator IdGenerator) ApiKeyService {
	return &apiKeyServiceImpl{
		dao:         dao,
		idGenerator: generator,
	}
}
