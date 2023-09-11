package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPasswordHasher(t *testing.T) {
	assert.Equal(t, NewPasswordHasher(), &passwordHasherImpl{})
}

func Test_passwordHasherImpl_Compare(t *testing.T) {
	hasher := &passwordHasherImpl{}
	hashed := hasher.Hash("plain_password")

	assert.NotEmpty(t, hashed)
}

func Test_passwordHasherImpl_Hash(t *testing.T) {
	hasher := &passwordHasherImpl{}
	hashed := hasher.Hash("plain_password")
	isTrue := hasher.Compare(*hashed, "plain_password")

	assert.True(t, isTrue)
}

func Test_passwordHasherImpl_Hash_False(t *testing.T) {
	hasher := &passwordHasherImpl{}
	hashed := hasher.Hash("plain_password")
	isTrue := hasher.Compare(*hashed, "false_password")
	assert.False(t, isTrue)
}
