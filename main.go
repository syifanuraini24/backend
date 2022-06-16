package main

import (
	"database/sql"
	"fmt"
	"log"
	"test/api"
	"test/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	fmt.Println("Database connected")

	userRepo := repository.NewUserRepository(db)

	mainAPI := api.NewAPI(*userRepo)
	mainAPI.Start()
}
