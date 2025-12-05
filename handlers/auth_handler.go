package handlers

import (
	"encoding/json"
	"net/http"
	"study/models"
	"study/storage"
	"study/utils"
)

type AuthHandler struct {
	userStorage *storage.UserStorage
}

func NewAuthHandler() *AuthHandler{
	return &AuthHandler{
		userStorage: storage.NewUserStorage(),
	}
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request){
	var req models.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w,  "invalid JSON", http.StatusBadRequest)
		return
	}

	user, err := a.userStorage.CreateUser(req.Email, req.Password)

	if err != nil{
		http.Error(w, "Error creating user ..." + err.Error(), http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateToken(user)
	if err != nil{
		http.Error(w, "Error generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
		"message": "User registered successfully",
	})



}


func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request){
	var req models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}


	user, err := a.userStorage.CheckUser(req.Email, req.Password)
	if err != nil{
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
	}

	token, err := utils.GenerateToken(user)
	if err != nil{
		http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "token": token,
        "message": "Login successful",
    })
}