package repository

import (
	"database/sql"
	"errors"
)

type AuthRepository interface {
	AddToken(refreshToken string) error
	DeleteToken(refreshToken string) error
	VerifyTokenIsExist(refreshToken string) error
}

type authRepositoryImpl struct {
	db *sql.DB
}

func (a *authRepositoryImpl) AddToken(refreshToken string) error {
	query := "INSERT INTO authentications (refresh_token) VALUES ($1)"
	_, err := a.db.Exec(query, refreshToken)

	return err
}

func (a *authRepositoryImpl) DeleteToken(refreshToken string) error {
	query := "DELETE FROM authentications WHERE refresh_token = $1"
	_, err := a.db.Exec(query, refreshToken)

	return err
}

func (a *authRepositoryImpl) VerifyTokenIsExist(refreshToken string) error {
	query := "SELECT * FROM authentications WHERE refresh_token = $1"
	rows, err := a.db.Query(query, refreshToken)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return nil
	}
	return errors.New("refreshToken does not exists")
}
