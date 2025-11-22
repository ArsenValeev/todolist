package storage

import (
	"fmt"
	"strconv"
	"strings"
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
        s.tasks[id].TaskDesc = description
        return s.tasks[id]
    }
    return nil // Задача не найдена
}

func (s *Storage) TaskCompleted(id int)*models.Task{
	if _, exist := s.tasks[id]; exist{
		s.tasks[id].Completed = true
		return s.tasks[id]
	} else{
		fmt.Println("Прозошла ошибка, не смогли обноваить стату задачи")
		return nil
	}
}


func (s *Storage) GetTaskWithFilter(completed, title string) []*models.Task{
	mapTasks := s.tasks 
	result := []*models.Task{}
	
	for _, task := range mapTasks{
		matches := false
		if completed != ""{
			comp, err := strconv.ParseBool(completed)
			if err != nil{
				fmt.Println("Ошибка конвертации в булевое значение")
			}
			if task.Completed == comp{
				matches = true
			}
		}

		if title != "" && strings.Contains(strings.ToLower(task.Title), strings.ToLower(title)){
			matches = true
		}

		if matches{
			result = append(result, task)
		}
	}
	return  result
}