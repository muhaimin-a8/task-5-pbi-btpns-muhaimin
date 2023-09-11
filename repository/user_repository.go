package repository

import (
	"database/sql"
	"errors"
	"pbi-btpns-api/entity"
	"time"
)

type UserRepository interface {
	VerifyUsernameNotExist(username string) error
	VerifyEmailNotExist(email string) error
	CreateUser(user entity.User) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUserById(userId string) (*entity.User, error)
	DeleteUserById(userId string) error
	UpdateUser(user entity.User) (*entity.User, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func (u *userRepositoryImpl) GetUserByEmail(email string) (*entity.User, error) {
	query := "SELECT * FROM users WHERE email = $1 AND is_deleted = false"
	rows, err := u.db.Query(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userEntity entity.User
	if rows.Next() {
		err := rows.Scan(&userEntity.Id, &userEntity.Username, &userEntity.Email, &userEntity.Password, &userEntity.IsDeleted, &userEntity.CreatedAt, &userEntity.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	if userEntity.Id == "" {
		return nil, errors.New("email not found")
	}

	return &userEntity, nil
}

func (u *userRepositoryImpl) GetUserById(userId string) (*entity.User, error) {
	query := "SELECT * FROM users WHERE id = $1 AND is_deleted = false"
	rows, err := u.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userEntity entity.User
	if rows.Next() {
		err := rows.Scan(&userEntity.Id, &userEntity.Username, &userEntity.Email, &userEntity.Password, &userEntity.IsDeleted, &userEntity.CreatedAt, &userEntity.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	if userEntity.Id == "" {
		return nil, errors.New("email not found")
	}

	return &userEntity, nil
}

func (u *userRepositoryImpl) DeleteUserById(userId string) error {
	// soft delete by updating `is_deleted` to `true`
	query := "UPDATE users SET is_deleted = true, updated_at = $1 WHERE id = $2 AND is_deleted = false"
	res, err := u.db.Exec(query, time.Now().Unix(), userId)

	if a, _ := res.RowsAffected(); a == 0 {
		return errors.New("userId not found")
	}
	return err
}

func (u *userRepositoryImpl) UpdateUser(user entity.User) (*entity.User, error) {
	query := "UPDATE users SET username = $1, email = $2, password = $3, updated_at = $4 WHERE id = $5 RETURNING id, username, email ,password ,is_deleted, created_at, updated_at"
	rows, err := u.db.Query(query, user.Username, user.Email, user.Password, user.UpdatedAt, user.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userEntity entity.User
	if rows.Next() {
		err := rows.Scan(&userEntity.Id, &userEntity.Username, &userEntity.Email, &userEntity.Password, &userEntity.IsDeleted, &userEntity.CreatedAt, &userEntity.UpdatedAt)

		if err != nil {
			return nil, err
		}
	}

	if userEntity.Id == "" {
		return nil, errors.New("userId not found")
	}
	return &userEntity, nil
}

func (u *userRepositoryImpl) VerifyUsernameNotExist(username string) error {
	query := "SELECT id FROM users WHERE username = $1"
	rows, err := u.db.Query(query, username)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		return errors.New("username already exists")
	}
	return nil
}

func (u *userRepositoryImpl) VerifyEmailNotExist(email string) error {
	query := "SELECT id FROM users WHERE email = $1"
	rows, err := u.db.Query(query, email)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		return errors.New("email already exists")
	}
	return nil
}

func (u *userRepositoryImpl) CreateUser(user entity.User) (*entity.User, error) {
	query := "INSERT INTO users(id, username, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, username, email ,password ,is_deleted, created_at, updated_at"
	rows, err := u.db.Query(query, user.Id, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userEntity entity.User
	if rows.Next() {
		err := rows.Scan(&userEntity.Id, &userEntity.Username, &userEntity.Email, &userEntity.Password, &userEntity.IsDeleted, &userEntity.CreatedAt, &userEntity.UpdatedAt)

		if err != nil {
			return nil, err
		}

	}

	return &userEntity, nil

}
