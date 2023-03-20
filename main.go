package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Lercc/go-gorm-restapi/db"
	"github.com/Lercc/go-gorm-restapi/env"
	"github.com/Lercc/go-gorm-restapi/models"
	"github.com/Lercc/go-gorm-restapi/routes"
	"github.com/rs/cors"
)

func main() {
	// Cargar variables
	env.LoadEnvironmentVariables()

	// Iniciar conexion a DB
	db.InitConnection()

	db.DBConnection.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.User{}, models.Task{})
	
	// Router
	router := routes.MuxRouter()

	//Cors
	ORIGINS := strings.Fields(os.Getenv("APP_CLIENT_ORIGINS"))
	routerCorsHandler := cors.New(cors.Options{ AllowedOrigins: ORIGINS }).Handler(router)

	PORT := os.Getenv("APP_PORT")
	log.Println("SERVER ON PORT :" + PORT)
	http.ListenAndServe(":" + PORT, routerCorsHandler)
}