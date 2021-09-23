package bootstrap

import (
	"crud/controllers"
	"crud/ws"

	"github.com/gorilla/mux"
)

// RouterInit function to create all routes
func RouterInit() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/user", controllers.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/user", controllers.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetPeopleEndpoint).Methods("GET")

	router.HandleFunc("/ws", ws.Serve)
	return router
}
