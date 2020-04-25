package routers

import (
	"net/http"

	"github.com/solrac87/rest/src/api/controllers"
)

var postRoutes = []Route{
	Route{
		Uri:     "/posts",
		Method:  http.MethodPost,
		Handler: controllers.CreatePost,
	},
}
