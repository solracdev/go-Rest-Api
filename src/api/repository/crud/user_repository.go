package crud

import (
	"context"
	"log"

	"github.com/solrac87/rest/src/api/database"
	"github.com/solrac87/rest/src/api/models"
	"github.com/solrac87/rest/src/api/utils/channels"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const usersCollectionName string = "users"
const counterCollectionName string = "counters"

type UserRepository struct {
	usersCollection   *mongo.Collection
	counterCollection *mongo.Collection
}

var User UserRepository

func (ur *UserRepository) Init(db *database.MongoDB) {
	ur.usersCollection = db.Database.Collection(usersCollectionName)
	ur.counterCollection = db.Database.Collection(counterCollectionName)
}

func (ur *UserRepository) Create(user models.User) (models.User, error) {

	done := make(chan bool)
	errorCh := make(chan error)

	go func(ch chan<- bool) {

		seq, err := ur.getNextSequence()
		if err != nil {
			ch <- false
			errorCh <- err
			return
		}

		user.ID = seq
		insertResult, err := ur.usersCollection.InsertOne(context.Background(), user)

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

		cursor, err := ur.usersCollection.Find(context.Background(), filter)

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

func (ur *UserRepository) FindById(id int64) (models.User, error) {

	filter := bson.M{"id": id}
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

	updateResult, err := ur.usersCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, err
	}

	return updateResult.ModifiedCount, nil
}

func (ur *UserRepository) Delete(id int) (int64, error) {

	filter := bson.M{"id": id}
	results, err := ur.usersCollection.DeleteOne(context.Background(), filter)

	if err != nil {
		return 0, err
	}

	return results.DeletedCount, nil
}

func (ur *UserRepository) getNextSequence() (int64, error) {

	filter := bson.M{"_id": usersCollectionName}
	update := bson.M{
		"$inc": bson.M{"seq": 1},
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := ur.counterCollection.FindOneAndUpdate(context.Background(), filter, update, &opt)

	if result.Err() != nil {
		return 0, result.Err()
	}

	doc := bson.M{}
	decodeErr := result.Decode(&doc)
	seq := doc["seq"].(int64)

	return seq, decodeErr
}
