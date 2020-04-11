package routers

import (
	"net/http"

	"github.com/solrac87/rest/src/api/controllers"
)

var usersRoutes = []Route{
	Route{
		Uri:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUsers,
	},
	Route{
		Uri:     "/users",
		Method:  http.MethodPost,
		Handler: controllers.CreateUser,
	},
	Route{
		Uri:     "/users/{nick}",
		Method:  http.MethodGet,
		Handler: controllers.GetUser,
	},
	Route{
		Uri:     "/users/{nick}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateUser,
	},
	Route{
		Uri:     "/users/{nick}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteUser,
	},
}
