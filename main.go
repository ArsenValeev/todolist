package main


import (
	"net/http"
	"study/handlers"
)

func main() {
	
	handlers.SetupRoutes()

	http.ListenAndServe(":8080", nil)

}




