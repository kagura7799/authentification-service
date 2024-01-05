package app

import (
	"authentification-service/internal/app/db"
	"fmt"
	"net/http"
)

type App struct {
	DB       *db.DB
}

func NewApp(db *db.DB) *App {
	return &App{
		DB:       db,
	}
}

func (a *App) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the auth project")
}

func (a *App) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Regist handler")

	username := r.FormValue("username")
	password := r.FormValue("password")

	a.DB.RegisterUser(username, password)

	fmt.Fprintf(w, "Registration successful for user %s", username)
}

func (a *App) LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login handler")

	username := r.FormValue("username")
	password := r.FormValue("password")
	
	isAuth, err := a.DB.AuthenticateUser(username, password)

	if isAuth {
		fmt.Fprintf(w, "Login successful.")
	} else {
		fmt.Fprintf(w, "Error, incorrect username or password")
	}

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
