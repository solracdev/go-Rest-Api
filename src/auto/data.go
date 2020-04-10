package auto

import (
	"time"

	"github.com/solrac87/rest/src/api/models"
)

var users = []models.User{
	models.User{
		NickName:  "Jhon Doe",
		Email:     "jhon@doe.com",
		Password:  "123456",
		CreatedAt: time.Now(),
	},
}
