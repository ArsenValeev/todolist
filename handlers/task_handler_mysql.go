package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"study/models"
	"study/storage"
)

type taskHandlerMySQL struct {
	db storage.MySQLStorage
}

func NewtaskHandlerMySQL() *taskHandlerMySQL{
	return &taskHandlerMySQL{
		db: *storage.NewMySQLStorage(),
	}
}

func (t *taskHandlerMySQL) CreateTaskDB(w http.ResponseWriter, r *http.Request) {
    userIDStr := r.Header.Get("User-ID")
    if userIDStr == "" {
        http.Error(w, "User not authenticated", http.StatusUnauthorized)
        return
    }

    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
		fmt.Println(userID, err)
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

   
    var req models.CreateTaskRequest
    err = json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }


    task := t.db.AddTask(req.Title, req.TaskDesc, userID)


    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func (t *taskHandlerMySQL) GetAllBD(w http.ResponseWriter, r *http.Request){
	userIDStr := r.Header.Get("User-ID")
	userID, _ := strconv.Atoi(userIDStr)



	tasks := t.db.GetTasksByUderID(userID)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)
} 

func (t *taskHandlerMySQL) GetTaskDBbyID(w http.ResponseWriter, r *http.Request){
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[len(parts) - 1])
	if err != nil{
		panic(err)
	}
	
	
	task := t.db.GetTaskByIDDB(id)

	if task == nil{
		fmt.Println("Задача не найдена")
	}

	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(task)
}

func (t *taskHandlerMySQL) DeleteTaskDB(w http.ResponseWriter, r *http.Request){
	str := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(str[len(str) - 1])
	if err != nil{
		panic(err)
	}
	t.db.DeleteTaskBD(id)
}

func (t *taskHandlerMySQL) UpdateDBbyID(w http.ResponseWriter, r *http.Request){
	str := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(str[len(str) - 1])
	if err != nil{
		panic(err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil{
		panic(err)
	}
	var stor models.CreateTaskRequest
	err = json.Unmarshal(body, &stor)

	if err != nil{
		panic(err)
	}

	
	task := t.db.UpdateByIDBD(stor.Title, stor.TaskDesc, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

