package repository

import (
	"github.com/stretchr/testify/assert"
	entity2 "pbi-btpns-api/internal/entity"
	"testing"
	"time"
)

func Test_photoRepositoryImpl_AddPhoto(t *testing.T) {
	defer photoTableTestHelper.CleanTable()
	defer userTableTestHelper.CleanTable()

	userTableTestHelper.CreateUser(entity2.User{
		Id: "user-123",
	})

	photoEntity := entity2.Photo{
		Id:        "photo-123",
		Title:     "John Doe",
		Caption:   "John Doe Profile Photo",
		Url:       "noisdfniaovnoiebvoadv",
		UserId:    "user-123",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	photo, err := photoRepository.AddPhoto(photoEntity)
	assert.Empty(t, err)
	assert.NotEmpty(t, photo)
	assert.Equal(t, photoEntity.Id, photo.Id)
	assert.Equal(t, photoEntity.Title, photo.Title)
	assert.Equal(t, photoEntity.Caption, photo.Caption)
	assert.Equal(t, photoEntity.Url, photo.Url)
	assert.Equal(t, photoEntity.UserId, photo.UserId)
	assert.Equal(t, photoEntity.CreatedAt, photo.CreatedAt)
	assert.Equal(t, photoEntity.UpdatedAt, photo.UpdatedAt)
}

func Test_photoRepositoryImpl_DeletePhotoById(t *testing.T) {
	defer photoTableTestHelper.CleanTable()
	defer userTableTestHelper.CleanTable()

	userTableTestHelper.CreateUser(entity2.User{
		Id: "user-123",
	})
	photoRepository.AddPhoto(entity2.Photo{
		Id:     "photo-123",
		UserId: "user-123",
	})

	err := photoRepository.DeletePhotoById("photo-123")
	assert.Empty(t, err)
}

func Test_photoRepositoryImpl_DeletePhotoById_IdNotFound(t *testing.T) {
	err := photoRepository.DeletePhotoById("photo-not-found")
	assert.NotEmpty(t, err)
}

func Test_photoRepositoryImpl_GetPhotoById(t *testing.T) {
	defer photoTableTestHelper.CleanTable()
	defer userTableTestHelper.CleanTable()

	userTableTestHelper.CreateUser(entity2.User{
		Id: "user-123",
	})

	photoEntity := entity2.Photo{
		Id:        "photo-123",
		Title:     "John Doe",
		Caption:   "John Doe Profile Photo",
		Url:       "noisdfniaovnoiebvoadv",
		UserId:    "user-123",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	photoTableTestHelper.AddPhoto(photoEntity)

	photo, err := photoRepository.GetPhotoById("photo-123")

	assert.Empty(t, err)
	assert.NotEmpty(t, photo)
	assert.Equal(t, photoEntity.Id, photo.Id)
	assert.Equal(t, photoEntity.Title, photo.Title)
	assert.Equal(t, photoEntity.Caption, photo.Caption)
	assert.Equal(t, photoEntity.Url, photo.Url)
	assert.Equal(t, photoEntity.UserId, photo.UserId)
	assert.Equal(t, photoEntity.CreatedAt, photo.CreatedAt)
	assert.Equal(t, photoEntity.UpdatedAt, photo.UpdatedAt)

}

func Test_photoRepositoryImpl_GetPhotoById_IdNotFound(t *testing.T) {

	photo, err := photoRepository.GetPhotoById("photo-not-found")

	assert.Empty(t, photo)
	assert.NotEmpty(t, err)

}
func Test_photoRepositoryImpl_UpdatePhoto(t *testing.T) {
	defer photoTableTestHelper.CleanTable()
	defer userTableTestHelper.CleanTable()

	userTableTestHelper.CreateUser(entity2.User{
		Id: "user-123",
	})

	photoTableTestHelper.AddPhoto(entity2.Photo{
		Id:        "photo-123",
		Title:     "John Doe",
		Caption:   "John Doe Profile Photo",
		Url:       "noisdfniaovnoiebvoadv",
		UserId:    "user-123",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})

	photoEntity := entity2.Photo{
		Id:        "photo-123",
		Title:     "John Doe updated",
		Caption:   "John Doe Profile Photo updated",
		Url:       "noisdfniaovnoiebvoadv-updated",
		UserId:    "user-123",
		UpdatedAt: time.Now().Add(time.Minute * 1).Unix(),
	}

	photo, err := photoRepository.UpdatePhoto(photoEntity)

	assert.Empty(t, err)
	assert.NotEmpty(t, photo)
	assert.Equal(t, photoEntity.Title, photo.Title)
	assert.Equal(t, photoEntity.Caption, photo.Caption)
	assert.Equal(t, photoEntity.Url, photo.Url)
	assert.Equal(t, photoEntity.UpdatedAt, photo.UpdatedAt)
	assert.NotEqual(t, photoEntity.CreatedAt, photoEntity.UpdatedAt)
}

func Test_photoRepositoryImpl_UpdatePhoto_IdNotFound(t *testing.T) {

	photoEntity := entity2.Photo{
		Id:        "photo-not-found",
		Title:     "John Doe updated",
		Caption:   "John Doe Profile Photo updated",
		Url:       "noisdfniaovnoiebvoadv-updated",
		UserId:    "user-123",
		UpdatedAt: time.Now().Add(time.Minute * 1).Unix(),
	}

	photo, err := photoRepository.UpdatePhoto(photoEntity)

	assert.Empty(t, photo)
	assert.NotEmpty(t, err)
}
