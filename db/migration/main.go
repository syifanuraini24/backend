package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		nama varchar(255), 
		email varchar(255), 
		password varhar(255),
		created_at datetime,
		updated_at datetime);`)
	if err != nil {
		log.Fatal("Error creating table: ", err)
	}
}
