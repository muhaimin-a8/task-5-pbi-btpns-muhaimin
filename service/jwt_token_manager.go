package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtTokenManager interface {
	NewAccessToken(userId string) (*string, error)
	NewRefreshToken(userId string) (*string, error)
	VerifyRefreshToken(token string) error
	ParseAccessToken(accessToken string) (*string, error)
	ParseRefreshToken(refreshToken string) (*string, error)
}

type jwtTokenManagerImpl struct {
	key             string
	expInSec        int
	refreshKey      string
	refreshExpInSec int
}

func (j *jwtTokenManagerImpl) VerifyRefreshToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(j.refreshKey), nil
	})

	if err != nil {
		return errors.New("token is expired")
	}
	return nil
}

func (j *jwtTokenManagerImpl) NewAccessToken(userId string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Second * time.Duration(j.expInSec)).Unix(),
		"sub": userId,
	})

	signedString, err := token.SignedString([]byte(j.key))
	if err != nil {
		return nil, err
	}

	return &signedString, nil
}

func (j *jwtTokenManagerImpl) NewRefreshToken(userId string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Second * time.Duration(j.refreshExpInSec)).Unix(),
		"sub": userId,
	})

	signedString, err := token.SignedString([]byte(j.refreshKey))
	if err != nil {
		return nil, err
	}

	return &signedString, nil
}

func (j *jwtTokenManagerImpl) ParseAccessToken(accessToken string) (*string, error) {
	parse, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(j.key), nil
	})

	if err != nil {
		return nil, err
	}
	sub, err := parse.Claims.GetSubject()
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

func (j *jwtTokenManagerImpl) ParseRefreshToken(refreshToken string) (*string, error) {
	parse, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(j.refreshKey), nil
	})

	if err != nil {
		return nil, err
	}
	sub, err := parse.Claims.GetSubject()
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

func NewJwtTokenManager(secret string, expInSec int, refreshSecret string, refreshExpInSec int) JwtTokenManager {
	return &jwtTokenManagerImpl{
		key:             secret,
		expInSec:        expInSec,
		refreshKey:      refreshSecret,
		refreshExpInSec: refreshExpInSec,
	}
}
