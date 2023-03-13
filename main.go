package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Lercc/go-gorm-restapi/db"
	"github.com/Lercc/go-gorm-restapi/env"
	"github.com/Lercc/go-gorm-restapi/routes"
)

func main() {
	// Cargar variables
	env.LoadEnvironmentVariables()

	// Iniciar conexion a DB
	db.InitConnection()
	
	// Router
	router := routes.MuxRouter()

	PORT := os.Getenv("APP_PORT")
	log.Println("SERVER ON PORT :" + PORT)
	http.ListenAndServe(":" + PORT, router)
}