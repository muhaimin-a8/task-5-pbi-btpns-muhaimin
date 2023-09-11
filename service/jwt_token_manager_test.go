package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_jwtTokenManagerImpl_NewAccessToken(t *testing.T) {
	jwtManager := &jwtTokenManagerImpl{
		key:             "key",
		expInSec:        90,
		refreshKey:      "key2",
		refreshExpInSec: 900,
	}

	token, err := jwtManager.NewAccessToken("user-123")
	assert.NotEmpty(t, token)
	assert.Empty(t, err)
}

func Test_jwtTokenManagerImpl_NewRefreshToken(t *testing.T) {
	jwtManager := &jwtTokenManagerImpl{
		key:             "key",
		expInSec:        90,
		refreshKey:      "key2",
		refreshExpInSec: 900,
	}

	token, err := jwtManager.NewRefreshToken("user-123")
	assert.NotEmpty(t, token)
	assert.Empty(t, err)
}

func Test_jwtTokenManagerImpl_Parse(t *testing.T) {
	jwtManager := &jwtTokenManagerImpl{
		key:             "key",
		expInSec:        10000,
		refreshKey:      "key2",
		refreshExpInSec: 900,
	}

	token, _ := jwtManager.NewAccessToken("user-123")

	userId, err := jwtManager.ParseAccessToken(*token)

	assert.Empty(t, err)
	assert.NotEmpty(t, userId)
	assert.Equal(t, *userId, "user-123")
}

func Test_jwtTokenManagerImpl_VerifyRefreshoken(t *testing.T) {
	jwtManager := &jwtTokenManagerImpl{
		key:             "key",
		expInSec:        9999,
		refreshKey:      "key2",
		refreshExpInSec: 110,
	}

	token, _ := jwtManager.NewRefreshToken("user-123")

	err := jwtManager.VerifyRefreshToken(*token)
	assert.Empty(t, err)

}

func Test_jwtTokenManagerImpl_VerifyRefreshoken_TokenExpired(t *testing.T) {
	jwtManager := &jwtTokenManagerImpl{
		key:             "key",
		expInSec:        9999,
		refreshKey:      "key2",
		refreshExpInSec: 0, // set expiration to 0
	}

	token, _ := jwtManager.NewRefreshToken("user-123")

	err := jwtManager.VerifyRefreshToken(*token)
	assert.NotEmpty(t, err)
	assert.Equal(t, errors.New("token is expired"), err)

}
