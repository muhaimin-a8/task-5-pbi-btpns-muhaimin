package test

import (
	"database/sql"
	"pbi-btpns-api/entity"
)

type PhotoTableTestHelper struct {
	db *sql.DB
}

func NewPhotoTableTestHelper(db *sql.DB) *PhotoTableTestHelper {
	return &PhotoTableTestHelper{db: db}
}

func (p *PhotoTableTestHelper) AddPhoto(photo entity.Photo) error {
	query := "INSERT INTO photos VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := p.db.Query(query, photo.Id, photo.Title, photo.Caption, photo.Url, photo.UserId, photo.CreatedAt, photo.UpdatedAt)

	return err
}

func (p *PhotoTableTestHelper) CleanTable() error {
	query := "DELETE FROM photos WHERE 1 = 1"
	_, err := p.db.Exec(query)

	return err
}
