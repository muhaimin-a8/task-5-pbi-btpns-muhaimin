package test

import (
	"database/sql"
)

type AuthTableTestHelper struct {
	db *sql.DB
}

func NewAuthTableTestHelper(db *sql.DB) *AuthTableTestHelper {
	return &AuthTableTestHelper{db: db}
}

func (a *AuthTableTestHelper) AddToken(refreshToken string) error {
	query := "INSERT INTO authentications (refresh_token) VALUES ($1)"
	_, err := a.db.Exec(query, refreshToken)

	return err
}

func (a *AuthTableTestHelper) IsExist(refreshToken string) (bool, error) {
	query := "SELECT * FROM authentications WHERE refresh_token = $1"
	rows, err := a.db.Query(query, refreshToken)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	var token string
	if rows.Next() {
		err := rows.Scan(&token)
		if err != nil {
			return false, err
		}
	}

	if token == "" {
		return false, nil
	}
	return true, nil
}

func (a *AuthTableTestHelper) CleanTable() error {
	query := "DELETE FROM authentications WHERE 1 = 1"
	_, err := a.db.Exec(query)

	return err
}
