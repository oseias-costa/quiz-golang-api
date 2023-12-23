package service

import (
	"testing"

	"github.com/oseias-costa/quiz-golang-api/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	var identifier int64 = 1
	var ageMock int64 = 32

	post := entity.Post{Id: 1, Name: "Oséias", Age: 32}
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)
	result, _ := testService.FindAll()

	mockRepo.AssertExpectations(t)
	assert.Equal(t, identifier, result[0].Id)
	assert.Equal(t, "Oséias", result[0].Name)
	assert.Equal(t, ageMock, result[0].Age)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	var ageMock int64 = 32
	post := entity.Post{Name: "Oséias", Age: 32}

	mockRepo.On("Save").Return(&post, nil)
	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)
	assert.NotNil(t, result.Id)
	assert.Equal(t, "Oséias", result.Name)
	assert.Equal(t, ageMock, result.Age)
	assert.Nil(t, err)
}

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
