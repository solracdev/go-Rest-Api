package services

import (
	"time"

	"github.com/solrac87/rest/src/api/models"
	"github.com/solrac87/rest/src/api/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type UserService struct {
	repository repository.UserRepositoryInterface
}

var User UserService

func (us *UserService) Init(r repository.UserRepositoryInterface) {
	us.repository = r
}

func (us *UserService) CreateUser(u models.User) (models.User, error) {

	u.Prepare()
	err := u.Validate("create")

	if err != nil {
		return models.User{}, err
	}

	user, err := us.repository.Create(u)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (us *UserService) FindAll() ([]models.User, error) {
	return us.repository.FindAll(bson.M{})
}

func (us *UserService) FindById(id int64) (models.User, error) {

	user, err := us.repository.FindById(id)

	if err != nil {
		return models.User{}, err
	}

	return user, err
}

func (us *UserService) Update(id int64, user models.User) (int64, error) {

	err := user.Validate("update")
	if err != nil {
		return 0, err
	}

	user.HashPassword(user.Password)
	filter := bson.M{"id": id}
	update := bson.D{
		{
			"$set", bson.D{
				{"nickname", user.NickName},
				{"email", user.Email},
				{"password", user.Password},
				{"updated_at", time.Now()},
			},
		},
	}

	return us.repository.Update(filter, update)
}

func (us *UserService) Delete(id int64) (int64, error) {
	return us.repository.Delete(id)
}
