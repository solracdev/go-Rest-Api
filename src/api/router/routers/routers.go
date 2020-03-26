package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func SetupRoutes(r *mux.Router) *mux.Router {
	return r
}
