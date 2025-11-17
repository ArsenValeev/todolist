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

type TaskHandler struct {
	storage storage.Storage
}

func NewTaskHandler() *TaskHandler{
	return &TaskHandler{
		storage: *storage.NewStorage(),
	}
}

func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil{
		fmt.Println("Ошибка чтеняи файла", err)
		return 
	}

	var request models.CreateTaskRequest
	err = json.Unmarshal(body, &request)
	if err != nil{
		fmt.Println("Ошибка чтеняи json", err)
		return 
	}

	task := t.storage.AddTask(request.Title, request.Description)
	
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(task)

}

func (t *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request){
	if r.URL.RawQuery != ""{
		completedFilter := r.URL.Query().Get("completed")
		search := r.URL.Query().Get("search")
	
		filteredTasks := t.storage.GetTaskWithFilter(completedFilter, search)
	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(filteredTasks)	
	} else{

		task := t.storage.GetAllTasks()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}

	

}

func (t *TaskHandler) GetTaskId(w http.ResponseWriter, r *http.Request){
	parts := strings.Split(r.URL.Path, "/")
	fmt.Println(r.URL.Path)
	idStr := parts[len(parts) - 1]

	id, err := strconv.Atoi(idStr)
	if err != nil{
		fmt.Println("Ошибка передачи id")
		return
	}

	task := t.storage.GetTaskByID(id)
	if task == nil{
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (t *TaskHandler) UpdateTaskById(w http.ResponseWriter, r *http.Request) {
    // Получаем ID из URL
    parts := strings.Split(r.URL.Path, "/")
    idStr := parts[len(parts)-1]
    
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Неверный ID", http.StatusBadRequest)
        return
    }

    // Читаем и парсим JSON
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Ошибка чтения", http.StatusBadRequest)
        return
    }

    var request models.CreateTaskRequest
    err = json.Unmarshal(body, &request)
    if err != nil {
        http.Error(w, "Неверный JSON", http.StatusBadRequest)
        return
    }

	fmt.Println(request.Title, request.Description)

    // ОБНОВЛЯЕМ задачу, а не создаем новую!
    task := t.storage.UpdateTask(id, request.Title, request.Description)
    if task == nil {
        http.Error(w, "Задача не найдена", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}


func (t *TaskHandler) TaskCompletedByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts) - 1]

	id, err := strconv.Atoi(idStr)
	if err != nil{
		fmt.Println("Ошибка чтения id")
	}

	task := t.storage.TaskCompleted(id)
	if task == nil{
		fmt.Println("Ошибка запсиси")
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)

	
}