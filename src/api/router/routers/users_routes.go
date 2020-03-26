package routers

import (
	"api/controllers"
	"net/http"
)

var usersRoutes = []Route{
	Route{
		Uri:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUser,
	},
}
