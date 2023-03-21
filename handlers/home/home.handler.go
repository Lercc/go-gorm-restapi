package home

import (
	"html/template"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("views/Home.html")

	if err != nil {
		log.Fatal("PAGE VIEW LOAD ERROR : ", err.Error())
	}
	view.Execute(w, nil)
}