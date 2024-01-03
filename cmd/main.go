package main

import (
	"authentification-service/internal/app"
	"fmt"
	"net/http"
)

func main() {
	myApp := app.NewApp()

	http.HandleFunc("/", myApp.HomeHandler)
	http.HandleFunc("/register", myApp.RegisterHandler)
	http.HandleFunc("/login", myApp.LoginHandler)

	fmt.Println("Сервер запущен на порту 8080...")
	http.ListenAndServe(":8080", nil)
}