package task

import (
	"encoding/json"
	"net/http"

	"github.com/Lercc/go-gorm-restapi/db"
	"github.com/Lercc/go-gorm-restapi/models"
	"github.com/gorilla/mux"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DBConnection.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DBConnection.Find(&task, params["task"])
	
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	json.NewEncoder(w).Encode(&task)
}

func StoreTask(w  http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	result :=  db.DBConnection.Create(&task)
	if err := result.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&task)
}

func UpdateTask(w  http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte("StoreTask:" + vars["task"]))
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DBConnection.Find(&task, params["task"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	db.DBConnection.Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}




