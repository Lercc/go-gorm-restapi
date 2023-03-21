package routes

import (
	"github.com/gorilla/mux"
	"github.com/Lercc/go-gorm-restapi/middleware"
	"github.com/Lercc/go-gorm-restapi/handlers/user"

)

func ApiMuxRoutes(router *mux.Router) {

	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.Use(middleware.ContentTypeApplicationJsonMiddleware)
	
	apiRouter.HandleFunc("/users", user.GetAllUsers).Methods("GET")
	apiRouter.HandleFunc("/users/{user}", user.GetUser).Methods("GET")
	apiRouter.HandleFunc("/users", user.StoreUser).Methods("POST")
	apiRouter.HandleFunc("/users/{user}", user.UpdateUser).Methods("PUT")
	apiRouter.HandleFunc("/users/{user}", user.DeleteUser).Methods("DELETE")
}
