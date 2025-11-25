package handlers

import "net/http"

func SetupRoutes() {
	// taskHandler := NewTaskHandler()
    taskHandlerMySQL := NewtaskHandlerMySQL()

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        taskHandlerMySQL.GetAllBD(w, r)
    case "POST":
        taskHandlerMySQL.CreateTaskDB(w, r)
    }
	})
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        taskHandlerMySQL.GetTaskDBbyID(w, r)
    case "DELETE":
        taskHandlerMySQL.DeleteTaskDB(w, r)
    case "PATCH":
        taskHandlerMySQL.CompletedTask(w, r)
    case "PUT":
        taskHandlerMySQL.UpdateDBbyID(w, r)
    }
})

}