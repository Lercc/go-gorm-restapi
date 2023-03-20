package routes

import (
	"github.com/gorilla/mux"
)

func MuxRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)
	
	router.HandleFunc("/users", GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{user}", GetUserHandler).Methods("GET")
	router.HandleFunc("/users", StoreUserHandler).Methods("POST")
	router.HandleFunc("/users/{user}", UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{user}", DeleteUsersHandler).Methods("DELETE")

	return router
}
