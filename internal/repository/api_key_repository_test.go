package repository

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_apiKeyRepositoryImpl_AddKey(t *testing.T) {
	defer apiKeyTableTestHelper.CleanTable()

	key := "key-123"
	err := apiKeyRepository.AddKey(key)
	assert.Empty(t, err)

	isExist := apiKeyTableTestHelper.IsExist(key)
	assert.True(t, isExist)
}

func Test_apiKeyRepositoryImpl_VerifyKeyIsExist(t *testing.T) {
	defer apiKeyTableTestHelper.CleanTable()

	key := "key-123"
	apiKeyTableTestHelper.AddKey(key)
	err := apiKeyRepository.VerifyKeyIsExist(key)
	assert.Empty(t, err)
}

func Test_apiKeyRepositoryImpl_VerifyKeyIsExist_KeyNotFound(t *testing.T) {
	err := apiKeyRepository.VerifyKeyIsExist("not-found-key")
	assert.NotEmpty(t, err)
	assert.Equal(t, errors.New("api key does not exists"), err)
}
