package service

import (
	"pbi-btpns-api/entity"
	"pbi-btpns-api/exception"
	"pbi-btpns-api/model"
	"pbi-btpns-api/repository"
	"time"
)

type PhotoService interface {
	AddPhoto(req model.AddPhotoRequest, userId string) *model.AddPhotoResponse
	GetPhotoById(photoId string, userId string) *model.GetPhotoResponse
	UpdatePhoto(req model.UpdatePhotoRequest, userId string) *model.UpdatePhotoResponse
	DeletePhoto(photoId string, userId string)
}

type photoServiceImpl struct {
	dao         repository.DAO
	idGenerator IdGenerator
}

func (p *photoServiceImpl) AddPhoto(req model.AddPhotoRequest, userId string) *model.AddPhotoResponse {
	photo := entity.Photo{
		Id:        "photo-" + p.idGenerator.New(20),
		Title:     req.Title,
		Caption:   req.Caption,
		Url:       req.ImgId,
		UserId:    userId,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	addedPhoto, err := p.dao.NewPhotoRepository().AddPhoto(photo)
	if err != nil {
		panic(exception.NotFoundError{Msg: "user id not found"})
	}

	return &model.AddPhotoResponse{
		Id:      addedPhoto.Id,
		Title:   addedPhoto.Title,
		Caption: addedPhoto.Caption,
		Url:     "/static/photos/" + addedPhoto.Url,
	}
}

func (p *photoServiceImpl) GetPhotoById(photoId string, userId string) *model.GetPhotoResponse {
	//TODO implement me
	panic("implement me")
}

func (p *photoServiceImpl) UpdatePhoto(req model.UpdatePhotoRequest, userId string) *model.UpdatePhotoResponse {
	//TODO implement me
	panic("implement me")
}

func (p *photoServiceImpl) DeletePhoto(photoId string, userId string) {
	//TODO implement me
	panic("implement me")
}

func NewPhotoService(dao repository.DAO, generator IdGenerator) PhotoService {
	return &photoServiceImpl{
		dao:         dao,
		idGenerator: generator,
	}
}
