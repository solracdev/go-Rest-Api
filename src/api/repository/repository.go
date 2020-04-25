package repository

import (
	"log"

	"github.com/solrac87/rest/src/api/database"
	"github.com/solrac87/rest/src/api/repository/crud"
)

// Load Repositories
func Load() {

	// Connect to mongo client Data Base
	var mongoDB database.MongoDB
	err := mongoDB.Connect()
	if err != nil {
		log.Fatal(err)
	}
	//defer mongoDB.Disconnect()

	// Initialize Uer Repo with data base connection
	crud.User.Init(&mongoDB)
	crud.Post.Init(&mongoDB)
}
