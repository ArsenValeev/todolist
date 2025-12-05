package storage

import (
	"database/sql"
	"study/models"

	"golang.org/x/crypto/bcrypt"
)

type UserStorage struct {
	db *sql.DB
}


func NewUserStorage() *UserStorage{
    db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/todolist?parseTime=true")
	if err != nil{
		panic(err)
	}

	return &UserStorage{
		db: db,
	}
}

func (u *UserStorage) CreateUser(email, password string) (*models.User, error){
	hasherPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return nil, err
	}

	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	result, err := u.db.Exec(query, email, string(hasherPassword))

	if err != nil{
		panic(err)
	}
	
	id, _ := result.LastInsertId()
	return &models.User{
		ID: int(id),
		Email: email,
	}, nil
}

func (u *UserStorage) CheckUser(email, password string)(*models.User, error){
    // 1. Ищем пользователя по email
    query := "SELECT id, email, password FROM users WHERE email = ?"
    row := u.db.QueryRow(query, email)
    
    var user models.User
    var hashedPassword string
    

    err := row.Scan(&user.ID, &user.Email, &hashedPassword)
    if err != nil {
        return nil, err
    }
    

    err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
        return nil, err 
    }
    

    return &user, nil
}
