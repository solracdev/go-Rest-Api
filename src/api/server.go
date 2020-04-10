package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/solrac87/rest/src/api/repository"
	"github.com/solrac87/rest/src/api/router"
	"github.com/solrac87/rest/src/api/services"
	"github.com/solrac87/rest/src/config"
)

func Run() {
	config.Load()
	repository.Load()
	services.Load()

	// auto.Load()
	fmt.Printf("\n\t Listening [::]:%d \n", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
