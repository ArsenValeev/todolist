package storage_test

import (
    "testing"
    "study/storage"
)
func TestAddTask(t *testing.T) {
	s := storage.NewMySQLStorage()

	task := s.AddTask("Test task", "Test description", 1)
	if task.Title != "Test task" {
        t.Errorf("Expected title 'Test task', got %s", task.Title)
    }
}

func TestGetByUserID(t *testing.T){
	s := storage.NewMySQLStorage()

	s.AddTask("Task 1", "Desc 1", 1)
    s.AddTask("Task 2", "Desc 2", 1)
    s.AddTask("Task 3", "Desc 3", 2)

	tasks := s.GetTasksByUderID(1)

	if len(tasks) != 3 {
        t.Errorf("Expected 2 tasks for user 1, got %d", len(tasks))
    }
}

func TestDeleteTask(t *testing.T) {
    s := storage.NewMySQLStorage()
    
    task := s.AddTask("To delete", "", 1)
    s.DeleteTaskBD(task.ID)
    
   
    
    taskAfter := s.GetTaskByIDDB(task.ID)
    if taskAfter != nil {
        t.Error("Task should not exist after deletion")
    }
}