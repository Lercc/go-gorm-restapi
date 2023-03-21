package routes

import (
	"github.com/gorilla/mux"
	"github.com/Lercc/go-gorm-restapi/handlers/home"
)

func WebMuxRoutes(router *mux.Router) {

	webRouter := router.NewRoute().Subrouter()
	webRouter.HandleFunc("/", home.HomePage)
}
