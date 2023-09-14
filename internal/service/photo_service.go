package service

import (
	"pbi-btpns-api/internal/entity"
	exception2 "pbi-btpns-api/internal/exception"
	"pbi-btpns-api/internal/model"
	"pbi-btpns-api/internal/repository"
	"time"
)

type PhotoService interface {
	AddPhoto(req model.AddPhotoRequest) *model.AddPhotoResponse
	GetPhotoById(photoId string, userId string) *model.GetPhotoResponse
	UpdatePhoto(req model.UpdatePhotoRequest) *model.UpdatePhotoResponse
	DeletePhoto(photoId string, userId string)
}

type photoServiceImpl struct {
	dao         repository.DAO
	idGenerator IdGenerator
}

func (p *photoServiceImpl) AddPhoto(req model.AddPhotoRequest) *model.AddPhotoResponse {
	photo := entity.Photo{
		Id:        "photo-" + p.idGenerator.New(20),
		Title:     req.Title,
		Caption:   req.Caption,
		Url:       req.Url,
		UserId:    req.UserId,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	addedPhoto, err := p.dao.NewPhotoRepository().AddPhoto(photo)
	if err != nil {
		panic(exception2.NotFoundError{Msg: "user id not found"})
	}

	return &model.AddPhotoResponse{
		Id:      addedPhoto.Id,
		Title:   addedPhoto.Title,
		Caption: addedPhoto.Caption,
		Url:     "/static/photos/" + addedPhoto.Url,
	}
}

func (p *photoServiceImpl) GetPhotoById(photoId string, userId string) *model.GetPhotoResponse {
	photo, err := p.dao.NewPhotoRepository().GetPhotoById(photoId)
	if err != nil {
		panic(exception2.NotFoundError{Msg: "photoId Not Found"})
	}

	// check owner
	if photo.UserId != userId {
		panic(exception2.AuthorizationError{Msg: "cannot get other photo"})
	}

	return &model.GetPhotoResponse{
		Id:      photo.Id,
		Title:   photo.Title,
		Caption: photo.Caption,
		Url:     "/static/photos/" + photo.Url,
	}
}

func (p *photoServiceImpl) UpdatePhoto(req model.UpdatePhotoRequest) *model.UpdatePhotoResponse {
	photoReq := entity.Photo{
		Id:        req.Id,
		Title:     req.Title,
		Caption:   req.Caption,
		Url:       req.Url,
		UserId:    req.UserId,
		UpdatedAt: time.Now().Unix(),
	}

	photo, err := p.dao.NewPhotoRepository().UpdatePhoto(photoReq)
	if err != nil {
		panic(exception2.NotFoundError{Msg: "photoId not found"})
	}

	return &model.UpdatePhotoResponse{
		Id:      photo.Id,
		Title:   photo.Title,
		Caption: photo.Caption,
		Url:     "/static/photos/" + photo.Url,
	}
}

func (p *photoServiceImpl) DeletePhoto(photoId string, userId string) {
	photo, err := p.dao.NewPhotoRepository().GetPhotoById(photoId)
	if err != nil {
		panic(exception2.NotFoundError{Msg: "photoId Not Found"})
	}

	// check owner
	if photo.UserId != userId {
		panic(exception2.AuthorizationError{Msg: "cannot delete other photo"})
	}

	err = p.dao.NewPhotoRepository().DeletePhotoById(photoId)
	if err != nil {
		panic(exception2.NotFoundError{Msg: "photoId Not Found"})
	}

	return
}

func NewPhotoService(dao repository.DAO, generator IdGenerator) PhotoService {
	return &photoServiceImpl{
		dao:         dao,
		idGenerator: generator,
	}
}
