package services

import (
	"github.com/solrac87/rest/src/api/models"
	"github.com/solrac87/rest/src/api/repository/crud"
)

type UserService struct {
	repository crud.UserRepository
}

func init() {

}

func (s *UserService) CreateUser(u models.User) {
	s.repository.Create(u)
}
