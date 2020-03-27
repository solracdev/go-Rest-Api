package api

import (
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	fmt.Println("\n\t Listening [::]:8085 \n")
	r := router.New()
	log.Fatal(http.ListenAndServe(":8085", r))
}
