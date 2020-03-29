package auto

import (
	"api/models"
	"time"
)

var users = []models.User{
	models.User{
		NickName:  "Jhon Doe",
		Email:     "jhon@doe.com",
		Password:  "123456",
		CreatedAt: time.Now(),
	},
}
