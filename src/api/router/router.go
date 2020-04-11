package router

import (
	"github.com/gorilla/mux"
	"github.com/solrac87/rest/src/api/router/routers"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routers.SetupRoutesWithMiddlewares(r)
}
