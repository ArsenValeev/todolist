package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"study/utils"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == ""{
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
            return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ParseToken(tokenString)
		if err != nil{
			http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
		}
		userIDStr := fmt.Sprintf("%d", claims.UserID)
		

		r.Header.Set("User-ID", userIDStr)
		next(w, r)
	}
}
