package database

import (
	"context"
	"log"
	"time"

	"github.com/solrac87/rest/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB struct
type MongoDB struct {
	ctx      context.Context
	client   *mongo.Client
	Database *mongo.Database
}

// Connect to DB
func (db *MongoDB) Connect() error {

	var err error
	db.client, err = mongo.NewClient(options.Client().ApplyURI(config.DBURL))

	if err != nil {
		log.Fatal(err)
	}

	db.ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = db.client.Connect(db.ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = db.client.Ping(db.ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	db.Database = db.client.Database(config.DBNAME)
	return err
}

// Disconnect from DB
func (db *MongoDB) Disconnect() {
	db.client.Disconnect(db.ctx)
}
