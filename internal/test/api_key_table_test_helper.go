package test

import (
	"database/sql"
)

type ApiKeyTableTestHelper struct {
	db *sql.DB
}

func NewApiKeyTableTestHelper(db *sql.DB) *ApiKeyTableTestHelper {
	return &ApiKeyTableTestHelper{db: db}
}

func (a *ApiKeyTableTestHelper) AddKey(key string) error {
	query := "INSERT INTO api_keys(key) VALUES ($1)"
	_, err := a.db.Exec(query, key)

	return err
}

func (a *ApiKeyTableTestHelper) IsExist(key string) bool {
	query := "SELECT * FROM api_keys WHERE key = $1"
	rows, err := a.db.Query(query, key)
	if err != nil {
		return false
	}
	defer rows.Close()

	if rows.Next() {
		return true
	}
	return false
}

func (a *ApiKeyTableTestHelper) CleanTable() error {
	query := "DELETE FROM api_keys WHERE 1 = 1"
	_, err := a.db.Exec(query)

	return err
}
