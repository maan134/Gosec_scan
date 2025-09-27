package main

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	db *sql.DB
)

func main() {
	db = initDB()
	defer db.Close()

	username := "admin' --"
	password := "any-password"

	err := authenticate(username, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Authentication successful")
}

func initDB() *sql.DB {
	var err error
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		password TEXT
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	insertUserQuery := `
	INSERT INTO users (username, password) VALUES ('admin', 'password123');
	`
	_, err = db.Exec(insertUserQuery)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func authenticate(username, password string) error {
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s';", username, password)
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return nil
	}

	return fmt.Errorf("invalid credentials")
}