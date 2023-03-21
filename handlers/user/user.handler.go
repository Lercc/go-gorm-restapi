package user

import (
	"encoding/json"
	"net/http"

	"github.com/Lercc/go-gorm-restapi/db"
	"github.com/Lercc/go-gorm-restapi/models"
	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DBConnection.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
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

func StoreUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	result := db.DBConnection.Create(&user)
	if err := result.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte("StoreUserHandler" + vars["user"]))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
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

