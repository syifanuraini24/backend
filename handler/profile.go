package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")

	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"password": password,
	})
}
