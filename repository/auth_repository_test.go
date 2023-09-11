package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_authRepositoryImpl_AddToken(t *testing.T) {
	defer authTableTestHelper.CleanTable()
	err := authRepository.AddToken("refreshToken")
	isExist, _ := authTableTestHelper.IsExist("refreshToken")

	assert.Empty(t, err)
	assert.True(t, isExist)
}

func Test_authRepositoryImpl_DeleteToken(t *testing.T) {
	defer authTableTestHelper.CleanTable()
	authTableTestHelper.AddToken("refreshToken")

	err := authRepository.DeleteToken("refreshToken")

	isExist, _ := authTableTestHelper.IsExist("refreshToken")

	assert.Empty(t, err)
	assert.False(t, isExist)
}

func Test_authRepositoryImpl_VerifyTokenIsExist(t *testing.T) {
	defer authTableTestHelper.CleanTable()
	authTableTestHelper.AddToken("refreshToken")

	err := authRepository.VerifyTokenIsExist("refreshToken")

	isExist, _ := authTableTestHelper.IsExist("refreshToken")

	assert.Empty(t, err)
	assert.True(t, isExist)
}
