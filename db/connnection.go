package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConnection: variable en la cual se guarda la conexión a la BD establecida
var DBConnection *gorm.DB

// InitConnection realiza la conexión con la DB y la guarda en una variable DBConnection, disponible desde el package db
func InitConnection() {

	fmt.Println("se llamo a la funcion InitConnection")
 
	// Data Source Name
	var DSN = os.Getenv("DB_DSN")
	
	var err error
	DBConnection, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal("BD CONNECCTION FAILED : ", err.Error())
	} else {
		log.Println("BD CONNECCTION SUCCESS")
	}
}