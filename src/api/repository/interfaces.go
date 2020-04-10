package repository

import "github.com/solrac87/rest/src/api/models"

// UserRepositoryInterface definition
type UserRepositoryInterface interface {
	Create(user models.User) (models.User, error)
	// FindAll() ([]models.User, error)
	// FindByNickname(n string) (models.User, error)
	// Update(n string, u models.User) (int32, error)
	// Delete(n string) (int32, error)
}
