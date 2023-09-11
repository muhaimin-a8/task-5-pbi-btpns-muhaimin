package service

import (
	"pbi-btpns-api/exception"
	"pbi-btpns-api/model"
	"pbi-btpns-api/repository"
)

type AuthService interface {
	Login(loginReq model.LoginRequest) *model.LoginResponse
	Logout(logoutReq model.LogoutRequest)
	UpdateToken(updateReq model.UpdateTokenRequest) *model.UpdateTokenResponse
}

type authServiceImpl struct {
	dao          repository.DAO
	hasher       PasswordHasher
	tokenManager JwtTokenManager
}

func (a *authServiceImpl) Login(loginReq model.LoginRequest) *model.LoginResponse {
	user, err := a.dao.NewUserRepository().GetUserByEmail(loginReq.Email)
	if err != nil {
		panic(exception.AuthenticationError{Msg: "invalid credential"})
	}

	match := a.hasher.Compare(user.Password, loginReq.Password)

	if !match {
		panic(exception.AuthenticationError{Msg: "invalid credential"})
	}

	// user is valid
	// generate accessToken and refreshToken
	accessToken, err := a.tokenManager.NewAccessToken(user.Id)
	if err != nil {
		panic(err)
	}

	refreshToken, err := a.tokenManager.NewRefreshToken(user.Id)
	if err != nil {
		panic(err)
	}

	// store refreshToken to database
	err = a.dao.NewAuthRepositroy().AddToken(*refreshToken)
	if err != nil {
		panic(err)
	}

	return &model.LoginResponse{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}
}

func (a *authServiceImpl) Logout(logoutReq model.LogoutRequest) {
	err := a.tokenManager.VerifyRefreshToken(logoutReq.RefreshToken)
	if err != nil {
		panic(exception.AuthenticationError{Msg: "invalid token"})
	}

	err = a.dao.NewAuthRepositroy().DeleteToken(logoutReq.RefreshToken)
	if err != nil {
		panic(err)
	}
}

func (a *authServiceImpl) UpdateToken(updateReq model.UpdateTokenRequest) *model.UpdateTokenResponse {
	err := a.tokenManager.VerifyRefreshToken(updateReq.RefreshToken)
	if err != nil {
		panic(exception.AuthenticationError{Msg: "invalid token"})
	}

	err = a.dao.NewAuthRepositroy().VerifyTokenIsExist(updateReq.RefreshToken)
	if err != nil {
		panic(exception.AuthenticationError{Msg: "invalid token"})
	}

	userId, err := a.tokenManager.ParseRefreshToken(updateReq.RefreshToken)
	if err != nil {
		panic(exception.AuthenticationError{Msg: "invalid token"})
	}

	newAccessToken, err := a.tokenManager.NewAccessToken(*userId)
	if err != nil {
		panic(err)
	}

	return &model.UpdateTokenResponse{
		AccessToken: *newAccessToken,
	}
}

func NewAuthService(dao repository.DAO, hasher PasswordHasher, tokenManager JwtTokenManager) AuthService {
	return &authServiceImpl{
		dao:          dao,
		hasher:       hasher,
		tokenManager: tokenManager,
	}
}
