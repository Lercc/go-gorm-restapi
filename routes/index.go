package routes

import (
	"github.com/gorilla/mux"
)

func MuxRouter() *mux.Router {
	router := mux.NewRouter()
	ApiMuxRoutes(router)
	WebMuxRoutes(router)
	return router
}
