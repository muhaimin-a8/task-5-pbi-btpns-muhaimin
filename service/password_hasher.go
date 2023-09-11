package service

import "golang.org/x/crypto/bcrypt"

type PasswordHasher interface {
	Hash(password string) *string
	Compare(hashed string, password string) bool
}

type passwordHasherImpl struct{}

func (p passwordHasherImpl) Hash(password string) *string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	hashedStr := string(hashed)

	return &hashedStr
}

func (p passwordHasherImpl) Compare(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func NewPasswordHasher() PasswordHasher {
	return &passwordHasherImpl{}
}
