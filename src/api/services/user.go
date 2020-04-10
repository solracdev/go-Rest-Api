package services

import (
	"api/models"
	"api/repository/crud"
)

type UserService struct {
	repository crud.UserRepository
}

func init() {

}

func (s *UserService) CreateUser(u models.User) {
	s.repository.Create(u)
}
