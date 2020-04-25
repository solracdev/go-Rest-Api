package services

import "github.com/solrac87/rest/src/api/repository/crud"

// Load Services
func Load() {

	// Initialize Services
	User.Init(&crud.User)
	Post.Init(&crud.Post)
}
