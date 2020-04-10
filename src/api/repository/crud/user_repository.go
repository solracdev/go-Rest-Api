package crud

import (
	"github.com/solrac87/rest/src/api/database"
	"github.com/solrac87/rest/src/api/models"
	"github.com/solrac87/rest/src/api/utils/channels"
	"go.mongodb.org/mongo-driver/mongo"
)

const collection string = "users"

type UserRepository struct {
	handle *mongo.Collection
}

func (ur *UserRepository) NewUserRepository(mc *database.MongoClient) *UserRepository {
	return &UserRepository{
		handle: mc.Database.Collection(collection),
	}
}

func (ur *UserRepository) Create(user models.User) (models.User, error) {

	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {

		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return user, nil
	}

	return models.User{}, err
}
