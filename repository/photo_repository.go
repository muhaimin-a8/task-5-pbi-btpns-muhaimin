package repository

import (
	"database/sql"
	"errors"
	"pbi-btpns-api/entity"
)

type PhotoRepository interface {
	AddPhoto(photo entity.Photo) (*entity.Photo, error)
	DeletePhotoById(photoId string) error
	UpdatePhoto(photo entity.Photo) (*entity.Photo, error)
	GetPhotoById(photoId string) (*entity.Photo, error)
}

type photoRepositoryImpl struct {
	db *sql.DB
}

func (p *photoRepositoryImpl) AddPhoto(photo entity.Photo) (*entity.Photo, error) {
	query := "INSERT INTO photos VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, title, caption, url, user_id, created_at, updated_at"
	rows, err := p.db.Query(query, photo.Id, photo.Title, photo.Caption, photo.Url, photo.UserId, photo.CreatedAt, photo.UpdatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photoEntity entity.Photo
	if rows.Next() {
		err := rows.Scan(&photoEntity.Id, &photoEntity.Title, &photoEntity.Caption, &photoEntity.Url, &photoEntity.UserId, &photoEntity.CreatedAt, &photoEntity.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &photoEntity, nil
}

func (p *photoRepositoryImpl) DeletePhotoById(photoId string) error {
	query := "DELETE FROM photos WHERE id = $1"
	res, err := p.db.Exec(query, photoId)

	if err != nil {
		return err
	}

	affected, _ := res.RowsAffected()
	if affected != 1 {
		return errors.New("photoId not found")
	}
	return nil
}

func (p *photoRepositoryImpl) UpdatePhoto(photo entity.Photo) (*entity.Photo, error) {
	query := "UPDATE photos SET title = $1, caption = $2, url = $3 , updated_at = $4 WHERE id = $5 RETURNING id, title, caption, url, user_id, created_at, updated_at"
	rows, err := p.db.Query(query, photo.Title, photo.Caption, photo.Url, photo.UpdatedAt, photo.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photoEntity entity.Photo
	if rows.Next() {
		err := rows.Scan(&photoEntity.Id, &photoEntity.Title, &photoEntity.Caption, &photoEntity.Url, &photoEntity.UserId, &photoEntity.CreatedAt, &photoEntity.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	if photoEntity.Id == "" {
		return nil, errors.New("photoId not found")
	}

	return &photoEntity, err
}

func (p *photoRepositoryImpl) GetPhotoById(photoId string) (*entity.Photo, error) {
	query := "SELECT * FROM photos WHERE id = $1"
	rows, err := p.db.Query(query, photoId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photoEntity entity.Photo
	if rows.Next() {
		err := rows.Scan(&photoEntity.Id, &photoEntity.Title, &photoEntity.Caption, &photoEntity.Url, &photoEntity.UserId, &photoEntity.CreatedAt, &photoEntity.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("photoId not found")
	}

	return &photoEntity, err
}
