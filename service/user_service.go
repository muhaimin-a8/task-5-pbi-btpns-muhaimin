package service

import (
	"pbi-btpns-api/entity"
	"pbi-btpns-api/exception"
	"pbi-btpns-api/model"
	"pbi-btpns-api/repository"
	"time"
)

type UserService interface {
	RegisterUser(registerModel model.UserRegisterRequest) *model.UserRegisterResponse
	UpdateUser(updateModel model.UserUpdateRequest) *model.UserUpdateResponse
	DeleteUserById(userId string)
}

type userServiceImpl struct {
	dao         repository.DAO
	idGenerator IdGenerator
	hasher      PasswordHasher
}

func (u *userServiceImpl) RegisterUser(registerModel model.UserRegisterRequest) *model.UserRegisterResponse {
	userEntity := entity.User{
		Id:        "user-" + u.idGenerator.New(20),
		Username:  registerModel.Username,
		Email:     registerModel.Email,
		Password:  *u.hasher.Hash(registerModel.Password),
		IsDeleted: false,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	err := u.dao.NewUserRepository().VerifyUsernameNotExist(registerModel.Username)
	if err != nil {
		panic(exception.InvariantError{Msg: "username already exist"})
	}

	err = u.dao.NewUserRepository().VerifyEmailNotExist(registerModel.Email)
	if err != nil {
		panic(exception.InvariantError{Msg: "email already exist"})
	}

	user, err := u.dao.NewUserRepository().CreateUser(userEntity)
	if err != nil {
		panic(err)
	}

	return &model.UserRegisterResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}

func (u *userServiceImpl) UpdateUser(updateModel model.UserUpdateRequest) *model.UserUpdateResponse {
	userEntity := entity.User{
		Id:        updateModel.Id,
		Username:  updateModel.Username,
		Email:     updateModel.Email,
		Password:  *u.hasher.Hash(updateModel.Password),
		IsDeleted: false,
		UpdatedAt: time.Now().Unix(),
	}

	userFromDB, err := u.dao.NewUserRepository().GetUserById(userEntity.Id)
	if err != nil {
		panic(exception.NotFoundError{Msg: "user id not found"})
	}

	// check
	if userFromDB.Username != userEntity.Username {
		err := u.dao.NewUserRepository().VerifyUsernameNotExist(updateModel.Username)
		if err != nil {
			panic(exception.InvariantError{Msg: "username already exist"})
		}

	}

	if userFromDB.Email != userEntity.Email {
		err = u.dao.NewUserRepository().VerifyEmailNotExist(updateModel.Email)
		if err != nil {
			panic(exception.InvariantError{Msg: "email already exist"})
		}

	}

	user, err := u.dao.NewUserRepository().UpdateUser(userEntity)
	if err != nil {
		panic(err)
	}

	return &model.UserUpdateResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}

func (u *userServiceImpl) DeleteUserById(userId string) {
	err := u.dao.NewUserRepository().DeleteUserById(userId)
	if err != nil {
		panic(exception.NotFoundError{Msg: "userId not found"})
	}
}

func NewUserService(dao repository.DAO, generator IdGenerator, hasher PasswordHasher) UserService {
	return &userServiceImpl{
		dao:         dao,
		idGenerator: generator,
		hasher:      hasher,
	}
}
