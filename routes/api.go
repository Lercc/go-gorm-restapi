package routes

import (
	"github.com/gorilla/mux"
	"github.com/Lercc/go-gorm-restapi/routes"
)

func MuxRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	return router
}