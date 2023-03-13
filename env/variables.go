package env

import (
	"log"
	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() {
	if err := godotenv.Load(".env"); err == nil {
		log.Println(".ENV LOAD SUCCESS")
	} else {
		log.Fatal(".ENV LOAD FAILED : ", err.Error())
	}
}