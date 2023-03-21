package task

import (
	"encoding/json"
	"net/http"

	"github.com/Lercc/go-gorm-restapi/db"
	"github.com/Lercc/go-gorm-restapi/models"
)

func getTasksController(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DBConnection.Find(&users)
	json.NewEncoder(w).Encode(&users)
}
