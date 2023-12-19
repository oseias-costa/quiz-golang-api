package service

import (
	"errors"
	"math/rand"

	"github.com/oseias-costa/quiz-golang-api/entity"
	"github.com/oseias-costa/quiz-golang-api/repository"
)

type PostService interface {
	Validate(p *entity.Post) error
	Create(p *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

func (*service) Validate(p *entity.Post) error {
	if p == nil {
		err := errors.New("post is empty")
		return err
	}
	if p.Name == "" {
		err := errors.New("Name is empty")
		return err
	}
	return nil

}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.Id = rand.Int63()
	return repo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
