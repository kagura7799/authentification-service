package main

import (
	"authentification-service/internal/app"
    "authentification-service/internal/app/db"
	"fmt"
	"net/http"
)

func main() {
	connStr := "user=kagura dbname=auth-db sslmode=disable"
    myDB, err := db.NewDB(connStr)
    if err != nil {
        panic(err)
    }
	
    defer myDB.Close()

    myApp := app.NewApp(myDB)

	http.HandleFunc("/", myApp.HomeHandler)
	http.HandleFunc("/register", myApp.RegisterHandler)
	http.HandleFunc("/login", myApp.LoginHandler)

	fmt.Println("Сервер запущен на порту 8080...")
	http.ListenAndServe(":8080", nil)
}