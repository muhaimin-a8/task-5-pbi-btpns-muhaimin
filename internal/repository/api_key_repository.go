package repository

import (
	"database/sql"
	"errors"
)

type ApiKeyRepository interface {
	AddKey(key string) error
	VerifyKeyIsExist(key string) error
}

type apiKeyRepositoryImpl struct {
	db *sql.DB
}

func (a *apiKeyRepositoryImpl) AddKey(key string) error {
	query := "INSERT INTO api_keys(key) VALUES ($1)"
	_, err := a.db.Exec(query, key)

	return err
}

func (a *apiKeyRepositoryImpl) VerifyKeyIsExist(key string) error {
	query := "SELECT * FROM api_keys WHERE key = $1"
	rows, err := a.db.Query(query, key)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return nil
	}
	return errors.New("api key does not exists")
}
