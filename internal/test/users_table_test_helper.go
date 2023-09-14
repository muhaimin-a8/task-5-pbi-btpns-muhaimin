package test

import (
	"database/sql"
	"pbi-btpns-api/internal/entity"
)

type UserTableTestHelper struct {
	db *sql.DB
}

func NewUserTableTestHelper(db *sql.DB) *UserTableTestHelper {
	return &UserTableTestHelper{db: db}
}

func (u *UserTableTestHelper) CreateUser(user entity.User) error {
	query := "INSERT INTO users(id, username, email, password,is_deleted, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := u.db.Exec(query, user.Id, user.Username, user.Email, user.Password, user.IsDeleted, user.CreatedAt, user.UpdatedAt)

	return err
}

func (u *UserTableTestHelper) CleanTable() error {
	query := "DELETE FROM users WHERE 1 = 1"
	_, err := u.db.Exec(query)

	return err
}
