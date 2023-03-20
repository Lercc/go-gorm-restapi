package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Lercc/go-gorm-restapi/db"
	"github.com/Lercc/go-gorm-restapi/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("GetUsersHandler"))
	var users []models.User
	db.DBConnection.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DBConnection.Find(&user, params["user"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func StoreUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	result := db.DBConnection.Create(&user)
	if err := result.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte("StoreUserHandler" + vars["user"]))
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DBConnection.Find(&user, params["user"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DBConnection.Delete(&user)
	// db.DBConnection.Unscoped().(&user)
	w.WriteHeader(http.StatusNoContent)
}
