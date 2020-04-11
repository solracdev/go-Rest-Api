package crud

import (
	"context"
	"log"

	"github.com/solrac87/rest/src/api/database"
	"github.com/solrac87/rest/src/api/models"
	"github.com/solrac87/rest/src/api/utils/channels"
	"go.mongodb.org/mongo-driver/bson"
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

		log.Println("Inserted ID -> ", insertResult)
		ch <- true
	}(done)

	if channels.OK(done) {
		return user, nil
	}

	return models.User{}, <-errorCh
}

func (ur *UserRepository) FindAll(filter interface{}) ([]models.User, error) {

	done := make(chan bool)
	errorCh := make(chan error)

	users := []models.User{}

	go func(ch chan<- bool) {

		cursor, err := ur.collection.Find(context.Background(), filter)

		if err != nil {
			ch <- false
			errorCh <- err
			return
		}

		for cursor.Next(context.Background()) {
			var user models.User
			err = cursor.Decode(&user)
			if err != nil {
				log.Fatal(err)
			}

			users = append(users, user)
		}

		cursor.Close(context.Background())
		ch <- true

	}(done)

	if channels.OK(done) {
		return users, nil
	}

	return users, <-errorCh
}

func (ur *UserRepository) FindByNickname(n string) (models.User, error) {

	filter := bson.M{"nickname": n}
	results, err := ur.FindAll(filter)

	if err != nil {
		return models.User{}, err
	}

	if len(results) == 0 {
		return models.User{}, nil
	}

	return results[0], nil
}

func (ur *UserRepository) Update(filter, update interface{}) (int64, error) {

	updateResult, err := ur.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, err
	}

	return updateResult.ModifiedCount, nil
}

func (ur *UserRepository) Delete(filter interface{}) (int64, error) {

	results, err := ur.collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return 0, err
	}

	return results.DeletedCount, nil
}
