package repository

import "database/sql"

type DAO interface {
	NewUserRepository() UserRepository
	NewAuthRepositroy() AuthRepository
	NewPhotoRepository() PhotoRepository
}

type daoImpl struct {
	db *sql.DB
}

func (d *daoImpl) NewPhotoRepository() PhotoRepository {
	return &photoRepositoryImpl{db: d.db}
}

func (d *daoImpl) NewUserRepository() UserRepository {
	return &userRepositoryImpl{db: d.db}
}

func (d *daoImpl) NewAuthRepositroy() AuthRepository {
	return &authRepositoryImpl{db: d.db}
}

func NewDAO(db *sql.DB) DAO {
	return &daoImpl{db: db}
}
