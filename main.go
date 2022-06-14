package main

import (
	"test/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.GET("/profile/:nama/", handler.Profile)

	r.Run(":8080")
}
