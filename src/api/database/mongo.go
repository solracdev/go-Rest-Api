package database

import (
	"context"
	"log"
	"time"

	"github.com/solrac87/rest/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDatabase interface
type MongoDatabase interface {
	Connect
	Disconnect
}

// MongoClient struct
type MongoClient struct {
	context  context.Context
	client   *mongo.Client
	Database *mongo.Database
}

// Connect to DB
func (mc *MongoClient) Connect() error {

	// Declare Context type object for managing multiple API requests
	mc.context, _ = context.WithTimeout(context.Background(), 15*time.Second)

	// Set client options
	clientOptions := options.Client().ApplyURI(config.DBURL)

	// Connect to MongoDB
	var err error
	mc.client, err = mongo.Connect(mc.context, clientOptions)

	// Check the connection
	err = mc.client.Ping(mc.context, nil)

	if err != nil {
		log.Fatal(err)
	}

	mc.Database = mc.client.Database(config.DBNAME)

	return nil
}

// Disconnect from DB
func (mc *MongoClient) Disconnect() {
	mc.client.Disconnect(mc.context)
}
