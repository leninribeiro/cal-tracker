package routes

import (
	"github.com/gorilla/mux"
	"github.com/leninribeiro/cal-tracker/controllers"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")

	return router
}
