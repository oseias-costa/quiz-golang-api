package service

import (
	"testing"

	"github.com/oseias-costa/quiz-golang-api/entity"
	"github.com/stretchr/testify/assert"
)

func TestPostService(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "post is empty", err.Error())
}

func TestValidadeEmptyName(t *testing.T) {
	post := entity.Post{Id: 1, Name: "", Age: 32}
	testService := NewPostService(nil)

	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "Name is empty", err.Error())
}
