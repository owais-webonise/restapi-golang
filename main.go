package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mirsaid-mirzohidov/restapi-golang/routes"
)

func main() {

	router := mux.NewRouter()
	routes.CreateRoute(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
