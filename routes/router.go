package routes

import (
	"github.com/gorilla/mux"
	"github.com/mirsaid-mirzohidov/restapi-golang/controllers"
)

func CreateRoute(router *mux.Router) {
	router.HandleFunc("/students", controllers.Getstudents).Methods("GET")
	router.HandleFunc("/students/{id}", controllers.GetStudent).Methods("GET")
	router.HandleFunc("/students", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/students/{id}", controllers.DeleteStudent).Methods("DELETE")

}
