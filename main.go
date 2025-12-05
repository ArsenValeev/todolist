package main


import (
	"net/http"
	"study/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)


	handlers.SetupRoutes()

	http.ListenAndServe(":8080", nil)

}





