package repository

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"log"
	app2 "pbi-btpns-api/internal/app"
	"pbi-btpns-api/internal/entity"
	test2 "pbi-btpns-api/internal/test"
	"testing"
	"time"
)

var userTableTestHelper *test2.UserTableTestHelper
var authTableTestHelper *test2.AuthTableTestHelper
var photoTableTestHelper *test2.PhotoTableTestHelper
var userRepository UserRepository
var authRepository AuthRepository
var photoRepository PhotoRepository

func TestMain(m *testing.M) {
	//load config
	err := app2.LoadConfig("../../config")
	if err != nil {
		log.Fatalln(err)
	}

	// setup db
	db, err := app2.NewDB()
	if err != nil {
		log.Fatalln(err)
	}

	// check connection to database
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	// setup test helper
	userTableTestHelper = test2.NewUserTableTestHelper(db)
	authTableTestHelper = test2.NewAuthTableTestHelper(db)
	photoTableTestHelper = test2.NewPhotoTableTestHelper(db)

	// setup repository instance
	userRepository = &userRepositoryImpl{db: db}
	authRepository = &authRepositoryImpl{db: db}
	photoRepository = &photoRepositoryImpl{db: db}

	// run all test cases
	m.Run()

	// close db connection
	err = db.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func Test_userRepositoryImpl_CreateUser(t *testing.T) {
	defer userTableTestHelper.CleanTable()

	userEntity := entity.User{
		Id:        "user-123",
		Username:  "johndoe",
		Email:     "johndoe@example.com",
		Password:  "johndoe123",
		IsDeleted: false,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	user, err := userRepository.CreateUser(userEntity)

	assert.Empty(t, err)
	assert.Equal(t, userEntity.Id, user.Id)
	assert.Equal(t, userEntity.Username, user.Username)
	assert.Equal(t, userEntity.Email, user.Email)
	assert.Equal(t, userEntity.Password, user.Password)
	assert.Equal(t, userEntity.IsDeleted, user.IsDeleted)
	assert.Equal(t, userEntity.CreatedAt, user.CreatedAt)
	assert.Equal(t, userEntity.UpdatedAt, user.UpdatedAt)
}

func Test_userRepositoryImpl_DeleteUserById(t *testing.T) {
	defer userTableTestHelper.CleanTable()

	user := entity.User{
		Id: "user-123",
	}
	userTableTestHelper.CreateUser(user)

	err := userRepository.DeleteUserById(user.Id)
	assert.Empty(t, err)
}

func Test_userRepositoryImpl_DeleteUserById_NotFoundUserId(t *testing.T) {
	defer userTableTestHelper.CleanTable()

	// must error when delete not found user
	err := userRepository.DeleteUserById("not-found-user-id")
	assert.NotEmpty(t, err, "must error but got empty")
	assert.Equal(t, errors.New("userId not found"), err)
}

func Test_userRepositoryImpl_GetUserByEmail(t *testing.T) {
	defer userTableTestHelper.CleanTable()

	user := entity.User{
		Id:        "user-123",
		Username:  "johndoe",
		Email:     "johndoe@example.com",
		Password:  "johndoe123",
		IsDeleted: false,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	userTableTestHelper.CreateUser(user)

	userEntity, err := userRepository.GetUserByEmail(user.Email)

	assert.Empty(t, err)
	assert.NotEmpty(t, userEntity)
	assert.Equal(t, userEntity.Id, user.Id)
	assert.Equal(t, userEntity.Username, user.Username)
	assert.Equal(t, userEntity.Email, user.Email)
	assert.Equal(t, userEntity.Password, user.Password)
	assert.Equal(t, userEntity.IsDeleted, user.IsDeleted)
	assert.Equal(t, userEntity.CreatedAt, user.CreatedAt)
	assert.Equal(t, userEntity.UpdatedAt, user.UpdatedAt)
}

func Test_userRepositoryImpl_GetUserByEmail_EmailNotFound(t *testing.T) {
	user, err := userRepository.GetUserByEmail("not-found-email")

	assert.Empty(t, user)
	assert.NotEmpty(t, err)
	assert.Equal(t, errors.New("email not found"), err)
}

func Test_userRepositoryImpl_UpdateUser(t *testing.T) {
	defer userTableTestHelper.CleanTable()
	beforeUpdate := entity.User{
		Id:        "user-123",
		Username:  "johndoe",
		Email:     "johndoe@example.com",
		Password:  "johndoe123",
		IsDeleted: false,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	updateRequest := entity.User{
		Id:        "user-123",
		Username:  "johndoeupdated",
		Email:     "johndoeupdated@example.com",
		Password:  "password updated",
		UpdatedAt: time.Now().Add(time.Minute * 1).Unix(),
	}

	userTableTestHelper.CreateUser(beforeUpdate)

	updatedUser, err := userRepository.UpdateUser(updateRequest)
	assert.Empty(t, err)
	assert.NotEmpty(t, updatedUser)
	assert.Equal(t, updateRequest.Username, updatedUser.Username)
	assert.Equal(t, updateRequest.Email, updatedUser.Email)
	assert.Equal(t, updateRequest.Password, updatedUser.Password)
	assert.Equal(t, updateRequest.UpdatedAt, updatedUser.UpdatedAt)
	assert.NotEqual(t, beforeUpdate.UpdatedAt, updatedUser.UpdatedAt)
}

func Test_userRepositoryImpl_UpdateUser_NotFoundUserId(t *testing.T) {
	updatedUser, err := userRepository.UpdateUser(entity.User{})
	assert.NotEmpty(t, err)
	assert.Equal(t, errors.New("userId not found"), err)
	assert.Empty(t, updatedUser)
}

func Test_userRepositoryImpl_VerifyEmailNotExist(t *testing.T) {
	err := userRepository.VerifyEmailNotExist("not-found-email")
	assert.Empty(t, err)
}

func Test_userRepositoryImpl_VerifyEmailNotExist_FoundEmail(t *testing.T) {
	defer userTableTestHelper.CleanTable()
	userTableTestHelper.CreateUser(entity.User{
		Email: "johndoe@example.com",
	})
	err := userRepository.VerifyEmailNotExist("johndoe@example.com")
	assert.NotEmpty(t, err)
	assert.Equal(t, errors.New("email already exists"), err)
}

func Test_userRepositoryImpl_VerifyUsernameNotExist(t *testing.T) {
	err := userRepository.VerifyUsernameNotExist("not-found-username")
	assert.Empty(t, err)
}

func Test_userRepositoryImpl_VerifyUsernameNotExist_FoundUsername(t *testing.T) {
	defer userTableTestHelper.CleanTable()
	userTableTestHelper.CreateUser(entity.User{
		Username: "johndoe",
	})
	err := userRepository.VerifyUsernameNotExist("johndoe")

	assert.NotEmpty(t, err)
	assert.Equal(t, errors.New("username already exists"), err)
}
