package main

import (
	"database/sql"
	"fmt"
	"log"
	"test/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	fmt.Println("Database connected")

	r := gin.Default()

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.GET("/profile/:nama/", handler.Profile)

	r.Run(":8080")
}
