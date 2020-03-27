package router

import (
	"api/router/routers"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routers.SetupRoutes(r)
}
