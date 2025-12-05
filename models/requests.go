package models

type CreateTaskRequest struct {
    Title       string `json:"title"`
    TaskDesc string `json:"task_desc"`
}

type RegisterRequest struct{
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginRequest struct{
    Email    string `json:"email"`
    Password string `json:"password"`
}