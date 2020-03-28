package database

import (
	"config"
	"context"
	"log"

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
	Client   *mongo.Client
	Database *mongo.Database
}

// Connect to DB
func (mc *MongoClient) Connect() error {

	// Set client options
	clientOptions := options.Client().ApplyURI(config.DBURL)

	// Connect to MongoDB
	var err error
	mc.Client, err = mongo.Connect(context.TODO(), clientOptions)

	// Check the connection
	err = mc.Client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	mc.Database = mc.Client.Database(config.DBNAME)
	return nil
}

// Disconnect from DB
func (mc *MongoClient) Disconnect() {
	defer mc.Client.Disconnect(context.TODO())
}
