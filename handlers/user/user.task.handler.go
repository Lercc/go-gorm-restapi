package user

import (
	"encoding/json"
	"net/http"

	"github.com/Lercc/go-gorm-restapi/db"
	"github.com/Lercc/go-gorm-restapi/models"
	"github.com/gorilla/mux"
)

func GetUserTasks(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var tasks []models.Task
	params := mux.Vars(r)
	db.DBConnection.Find(&user, params["user"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DBConnection.Where("user_id = ?", user.ID).Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}