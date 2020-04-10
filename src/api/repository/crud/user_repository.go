package crud

import (
	"context"
	"log"

	"github.com/solrac87/rest/src/api/database"
	"github.com/solrac87/rest/src/api/models"
	"github.com/solrac87/rest/src/api/utils/channels"
	"go.mongodb.org/mongo-driver/mongo"
)

const collection string = "users"

type UserRepository struct {
	collection *mongo.Collection
}

var User UserRepository

func (ur *UserRepository) Init(db *database.MongoDB) {
	ur.collection = db.Database.Collection(collection)
}

func (ur *UserRepository) Create(user models.User) (models.User, error) {

	done := make(chan bool)
	errorCh := make(chan error)

	go func(ch chan<- bool) {

		insertResult, err := ur.collection.InsertOne(context.Background(), user)

		if err != nil {
			ch <- false
			errorCh <- err
			return
		}

		log.Println("Inserted count -> ", insertResult)
		ch <- true
	}(done)

	if channels.OK(done) {
		return user, nil
	}

	return models.User{}, <-errorCh
}
