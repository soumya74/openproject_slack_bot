package main

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/slackBot/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	portalRouter := router.PathPrefix("/").Subrouter()
	handlers.HandleRoutes(portalRouter)
	fmt.Println("listening at Port 9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
