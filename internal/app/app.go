package app

import (
	"fmt"
	"net/http"
)

type App struct {
	Users    []User
	Sessions map[string]*Session
}

func NewApp() *App {
	return &App{
		Users:    make([]User, 0),
		Sessions: make(map[string]*Session),
	}
}

func (a *App) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the auth project")
}

func (a *App) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Regist handler")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Username and password are required.", http.StatusBadRequest)
		return
	}

	for _, user := range a.Users {
		if user.Username == username {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}
	}

	newUser := User{
		ID:       len(a.Users) + 1,
		Username: username,
		Password: password,
	}

	a.Users = append(a.Users, newUser)

	fmt.Fprint(w, "Registration successful for user %s", username)

}

func (a *App) LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Login handler")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Username and password are required.", http.StatusBadRequest)
		return
	}

	var foundUser User

	for _, user := range a.Users {
		if user.Username == username && user.Password == password {
			foundUser = user
			break
		}
	}

	if foundUser.ID == 0 {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}

	token := fmt.Sprintf("token_%d", foundUser.ID)
	session := &Session{UserID: foundUser.ID, Token: token}
	a.Sessions[token] = session

	fmt.Fprintf(w, "Login successful. Session token: %s", token)
}
