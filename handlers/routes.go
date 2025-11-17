package handlers

import "net/http"

func SetupRoutes() {
	taskHandler := NewTaskHandler()

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        taskHandler.GetTasks(w, r)
    case "POST":
        taskHandler.CreateTask(w, r)
    }
	})
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        taskHandler.GetTaskId(w, r)
    case "PUT":
        taskHandler.UpdateTaskById(w, r)
    }
})

}