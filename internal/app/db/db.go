package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const ConnStr string = "user=kagura dbname=auth-db sslmode=disable"

type DB struct {
	conn *sql.DB
}

func NewDB(connStr string) (*DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &DB{conn: db}, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

func (db *DB) Ping() error {
	return db.conn.Ping()
}

func UsageBD() {
	db, err := NewDB(ConnStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database!")
}

func (db *DB) isUsernameTaken(username string) (bool, error) {
	var count int
	err := db.conn.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1", username).Scan(&count)
	
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (db *DB) RegisterUser(username, password string) error {
	isTaken, err := db.isUsernameTaken(username)

	if err != nil {
		return err
	}

	if isTaken {
		return fmt.Errorf("Username %s already is taken", username)
	}

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    _, err = db.conn.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)
    return err
}

func (db *DB) AuthenticateUser(username, password string) (bool, error) {
    var hashedPassword string
    err := db.conn.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&hashedPassword)
    if err != nil {
        if err == sql.ErrNoRows {
            return false, nil	
        }
        return false, err
    }

    err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil, nil
}
