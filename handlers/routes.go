package handlers

import (
	"net/http"
	"study/middleware"
)

func SetupRoutes() {
	// taskHandler := NewTaskHandler()
    taskHandlerMySQL := NewtaskHandlerMySQL()

    AuthHandler := NewAuthHandler()


    //вход и регистрация
    http.HandleFunc("/register", AuthHandler.Register)
    http.HandleFunc("/login", AuthHandler.Login)



	http.HandleFunc("/tasks", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            taskHandlerMySQL.GetAllBD(w, r)
        case "POST":
            taskHandlerMySQL.CreateTaskDB(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    }))

    http.HandleFunc("/tasks/", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            taskHandlerMySQL.GetTaskDBbyID(w, r)
        case "DELETE":
            taskHandlerMySQL.DeleteTaskDB(w, r)
        case "PATCH":
            taskHandlerMySQL.CompletedTask(w, r)
        case "PUT":
            taskHandlerMySQL.UpdateDBbyID(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    }))

}