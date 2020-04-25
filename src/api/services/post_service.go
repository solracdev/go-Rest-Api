package services

import (
	"github.com/solrac87/rest/src/api/models"
	"github.com/solrac87/rest/src/api/repository"
)

type PostService struct {
	repository repository.PostRepositoryInterface
}

var Post PostService

func (ps *PostService) Init(r repository.PostRepositoryInterface) {
	ps.repository = r
}

func (ps *PostService) CreatePost(p models.Post) (models.Post, error) {

	p.Prepare()
	err := p.Valitade()
	if err != nil {
		return models.Post{}, err
	}

	post, err := ps.repository.Create(p)
	if err != nil {
		return models.Post{}, err
	}

	return post, err
}
