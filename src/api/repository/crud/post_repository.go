package crud

import (
	"context"
	"log"

	"github.com/solrac87/rest/src/api/models"
	"github.com/solrac87/rest/src/api/utils/channels"

	"github.com/solrac87/rest/src/api/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const postCollectionName string = "posts"

type PostRepository struct {
	postCollection    *mongo.Collection
	counterCollection *mongo.Collection
}

var Post PostRepository

func (pr *PostRepository) Init(db *database.MongoDB) {
	pr.postCollection = db.Database.Collection(postCollectionName)
	pr.counterCollection = db.Database.Collection(counterCollectionName)
}

func (pr *PostRepository) Create(post models.Post) (models.Post, error) {
	done := make(chan bool)
	errorCh := make(chan error)

	go func(ch chan<- bool) {

		seq, err := pr.getNextSequence()
		if err != nil {
			ch <- false
			errorCh <- err
			return
		}

		post.ID = seq
		insertResult, err := pr.postCollection.InsertOne(context.Background(), post)

		if err != nil {
			ch <- false
			errorCh <- err
			return
		}

		log.Println("Inserted ID -> ", insertResult)
		ch <- true
	}(done)

	if channels.OK(done) {
		return post, nil
	}

	return models.Post{}, <-errorCh
}

func (pr *PostRepository) getNextSequence() (int64, error) {

	filter := bson.M{"_id": postCollectionName}
	update := bson.M{
		"$inc": bson.M{"seq": 1},
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := pr.counterCollection.FindOneAndUpdate(context.Background(), filter, update, &opt)
	if result.Err() != nil {
		return 0, result.Err()
	}

	doc := bson.M{}
	decodeErr := result.Decode(&doc)
	seq := int64(doc["seq"].(int32))

	return seq, decodeErr
}
