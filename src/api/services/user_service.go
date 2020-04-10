package services

import (
	"log"

	"github.com/solrac87/rest/src/api/models"
	"github.com/solrac87/rest/src/api/repository"
)

type UserService struct {
	repository repository.UserRepositoryInterface
}

var User UserService

func (us *UserService) Init(r repository.UserRepositoryInterface) {
	us.repository = r
}

func (us *UserService) CreateUser(u models.User) (models.User, error) {

	u.HashPassword(u.Password)
	user, err := us.repository.Create(u)

	if err != nil {
		log.Printf("aSDS")
		return user, err
	}

	return user, nil
}
