package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewIdGenerator(t *testing.T) {
	assert.Equal(t, NewIdGenerator(), &idGeneratorImpl{})
}

func Test_idGeneratorImpl_New(t *testing.T) {
	generator := &idGeneratorImpl{}

	id := generator.New(20)
	assert.NotEmpty(t, id)
	assert.Len(t, id, 20)
}
