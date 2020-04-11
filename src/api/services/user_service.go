package services

import (
	"errors"
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

	u.HashPassword(u.Password)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	user, err := us.repository.Create(u)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (us *UserService) FindAll() ([]models.User, error) {
	return us.repository.FindAll(bson.M{})
}

func (us *UserService) FindByNickname(n string) (models.User, error) {

	user, err := us.repository.FindByNickname(n)

	if err != nil {
		return models.User{}, err
	}

	if user.NickName == "" {
		return models.User{}, errors.New("User not found with nickname: " + n)
	}

	return user, err
}

func (us *UserService) Update(n string, user models.User) (int64, error) {

	// find user by URL param nickname
	_, err := us.FindByNickname(n)

	if err != nil {
		return 0, err
	}

	// find user by body post nickname
	existUser, _ := us.FindByNickname(user.NickName)

	if (existUser != models.User{}) {
		return 0, errors.New("User already exist with nickname: " + n)
	}

	user.HashPassword(user.Password)
	filter := bson.M{"nickname": n}
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

func (us *UserService) Delete(n string) (int64, error) {
	filter := bson.M{"nickname": n}
	return us.repository.Delete(filter)
}
