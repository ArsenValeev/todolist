package models

type CreateTaskRequest struct {
    Title       string `json:"title"`
    TaskDesc string `json:"task_desc"`
}