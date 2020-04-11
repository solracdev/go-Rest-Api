package repository

import "github.com/solrac87/rest/src/api/models"

// UserRepositoryInterface definition
type UserRepositoryInterface interface {
	Create(user models.User) (models.User, error)
	FindAll(filter interface{}) ([]models.User, error)
	FindByNickname(n string) (models.User, error)
	Update(filter, update interface{}) (int64, error)
	Delete(filter interface{}) (int64, error)
}
