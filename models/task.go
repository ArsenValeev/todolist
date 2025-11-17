package models

import "time"

type Task struct {
	ID          int `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
}


func NewTask(title, description string) *Task{
	return &Task{
		Title: title,
		Description: description,
		CreatedAt: time.Now(),
	}
}

func (t *Task) MarkCompleted(){
	t.Completed = true
}

func (t *Task) MarkIncomplete(){
	t.Completed = false
}

