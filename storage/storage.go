package storage

import (
	"fmt"
	"study/models"
)


type Storage struct {
	nextId int
	tasks map[int]*models.Task
}

func NewStorage() *Storage{
	return &Storage{
		tasks: map[int]*models.Task{},
	}
}


func (s *Storage) AddTask(title, description string) *models.Task {
	result := models.NewTask(title, description)
	result.ID = s.nextId
	s.nextId++
	s.tasks[result.ID] = result
	return result
}

func (s *Storage) GetAllTasks() []*models.Task{
	result := []*models.Task{}
	for _, v := range s.tasks{
		result = append(result, v)
	}
	return result
}
func (s *Storage) GetTaskByID(id int) *models.Task{
	if _, ok := s.tasks[id]; !ok{
		fmt.Println("Неверный ID пользователя!")
		return nil
	}
	return s.tasks[id]
}


func (s *Storage) DeleteTask(id int) bool{
	if _, ok := s.tasks[id]; !ok{
		fmt.Println("Неверный ID пользователя!")
		return false
	}
	delete(s.tasks, id)
	return true
}


func (s *Storage) UpdateTask(id int, title, description string) *models.Task {
    // Проверяем существует ли задача
    if _, exists := s.tasks[id]; exists {
        // Обновляем поля
        s.tasks[id].Title = title
        s.tasks[id].Description = description
        return s.tasks[id]
    }
    return nil // Задача не найдена
}

func (s *Storage) TaskCompleted(id int)*models.Task{
	if task, exist := s.tasks[id]; exist{
		task.Completed = true
		return task
	} else{
		fmt.Println("Прозошла ошибка, не смогли обноваить стату задачи")
		return nil
	}
}

