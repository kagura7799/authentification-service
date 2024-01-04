package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

// NewDB создает новый объект базы данных.
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
	connStr := "user=kagura dbname=auth-db sslmode=disable"
	db, err := NewDB(connStr)
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
