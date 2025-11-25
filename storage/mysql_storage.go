package storage

import (
	"database/sql"
	"study/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
    db *sql.DB
}

func NewMySQLStorage() *MySQLStorage {

    db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/todolist?parseTime=true")
    if err != nil {
        panic(err)
    }
    return &MySQLStorage{db: db}
}


func (m *MySQLStorage) AddTask(title, taskDesc string) *models.Task{
	query := "INSERT INTO tasks (title, task_desc) VALUES (?, ?)"

	result, err := m.db.Exec(query, title, taskDesc)

	if err != nil{
		panic(err)

	}

	id, _ := result.LastInsertId()

	return &models.Task{
		ID: int(id),
		Title: title,
		TaskDesc: taskDesc,
		Completed: false,
		CreatedAt: time.Now(),
	}
}

func (m *MySQLStorage) GetAllTaskBD() []*models.Task {
    query := "SELECT * FROM tasks"
    rows, err := m.db.Query(query)
    if err != nil {
        panic(err)
    }
    defer rows.Close()  

    var tasks []*models.Task
    
    for rows.Next() {
        var task models.Task
        err := rows.Scan(
            &task.ID,
            &task.Title, 
            &task.TaskDesc,
            &task.Completed,
            &task.CreatedAt,
        )
        if err != nil {
            panic(err)
        }
        tasks = append(tasks, &task)
    }
    
    return tasks
}

func (m *MySQLStorage) GetTaskByIDDB(id int) *models.Task{
	query := "SELECT * FROM tasks WHERE id=?"
	row, err := m.db.Query(query, id)
	if err != nil{
		panic(err)
	}
	defer row.Close()

	if !row.Next(){
		return nil
	}
	task := &models.Task{}

	err = row.Scan(
            &task.ID,
            &task.Title, 
            &task.TaskDesc,
            &task.Completed,
            &task.CreatedAt,
        )
	if err != nil{
		panic(err)
	}
	return task
}

func (m *MySQLStorage) DeleteTaskBD(id int){
	quary := "DELETE FROM tasks WHERE id = ?"
	m.db.Exec(quary, id)
}


func (m *MySQLStorage) UpdateByIDBD(title, task_desk string, id int) *models.Task{

	cur := m.GetTaskByIDDB(id)
	if cur == nil{
		return nil
	}

	query := "UPDATE tasks SET title = ?, task_desc = ? WHERE id = ?"
	_, err := m.db.Exec(query, title, task_desk, id)
	if err != nil{
		panic(err)
	}
	
	
	return &models.Task{
		ID: id,
		Title: title,
		TaskDesc: task_desk,
		Completed: cur.Completed,
		CreatedAt: cur.CreatedAt,
	}
}


func (m *MySQLStorage) CompletedbyID(id int) *models.Task{
	quary := "UPDATE tasks SET Completed = true WHERE id = ?"


	cur := m.GetTaskByIDDB(id)
	if cur == nil{
		return nil
	}
	_, err := m.db.Exec(quary, id)
	if err != nil{
		panic(err)
	}

	return &models.Task{
		ID: cur.ID,
		Title: cur.Title,
		TaskDesc: cur.TaskDesc,
		Completed: true,
		CreatedAt: cur.CreatedAt,
	}
}