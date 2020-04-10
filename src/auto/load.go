package auto

import (
	"context"
	"fmt"
	"log"

	"github.com/solrac87/rest/src/api/database"
	"github.com/solrac87/rest/src/api/utils/console"
)

// Load Data into DB
func Load() {

	var mongoClient database.MongoClient
	err := mongoClient.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect()

	collection := mongoClient.Database.Collection("users")

	for _, u := range users {
		u.HashPassword(u.Password)
		fmt.Println(u.CreatedAt)
		insertResult, err := collection.InsertOne(context.TODO(), u)

		if err != nil {
			log.Fatal(err)
		}

		console.Pretty(insertResult)
	}
}
