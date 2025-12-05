package models

import (
	"database/sql"
	"time"
)

type Task struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    TaskDesc  string    `json:"taskDesc"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"createdAt"`
    UserID    sql.NullInt64 `json:"userId"`
}


func NewTask(title, description string) *Task{
	return &Task{
		Title: title,
		TaskDesc: description,
		CreatedAt: time.Now(),
	}
}

func (t *Task) MarkCompleted(){
	t.Completed = true
}

func (t *Task) MarkIncomplete(){
	t.Completed = false
}

