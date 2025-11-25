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

func (t *taskHandlerMySQL) CreateTaskDB(w http.ResponseWriter, r *http.Request){
	body, err := io.ReadAll(r.Body)
	if err != nil{
		fmt.Println("Ошибка чтения запроса")
	}


	var request models.CreateTaskRequest
	err = json.Unmarshal(body, &request)
	if err != nil{
		fmt.Println("Ошибка чтеняи json", err)
		return 
	}
	
	task := t.db.AddTask(request.Title, request.TaskDesc)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(task)

}

func (t *taskHandlerMySQL) GetAllBD(w http.ResponseWriter, r *http.Request){

	tasks := t.db.GetAllTaskBD()

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

