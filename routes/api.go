package routes

import (
	"github.com/gorilla/mux"
)

func MuxRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)

	return router
}