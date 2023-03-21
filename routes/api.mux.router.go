package routes

import (
	"github.com/gorilla/mux"
	"github.com/Lercc/go-gorm-restapi/middleware"
	"github.com/Lercc/go-gorm-restapi/handlers/user"
	"github.com/Lercc/go-gorm-restapi/handlers/task"

)

func ApiMuxRoutes(router *mux.Router) {

	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.Use(middleware.ContentTypeApplicationJsonMiddleware)
	
	// USER
	apiRouter.HandleFunc("/users", user.GetAllUsers).Methods("GET")
	apiRouter.HandleFunc("/users/{user}", user.GetUser).Methods("GET")
	apiRouter.HandleFunc("/users", user.StoreUser).Methods("POST")
	apiRouter.HandleFunc("/users/{user}", user.UpdateUser).Methods("PUT")
	apiRouter.HandleFunc("/users/{user}", user.DeleteUser).Methods("DELETE")

	// TASK
	apiRouter.HandleFunc("/tasks", task.GetAllTasks).Methods("GET")
	apiRouter.HandleFunc("/tasks/{task}", task.GetTask).Methods("GET")
	apiRouter.HandleFunc("/tasks", task.StoreTask).Methods("POST")
	apiRouter.HandleFunc("/tasks/{task}", task.UpdateTask).Methods("PUT")
	apiRouter.HandleFunc("/tasks/{task}", task.DeleteTask).Methods("DELETE")
}
