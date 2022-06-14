package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	nama := c.Param("nama")

	c.JSON(http.StatusOK, gin.H{
		"nama": nama,
	})
}
