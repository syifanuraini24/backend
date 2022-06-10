package handler

import (
	"fmt"
	"net/http"

	"test/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"password": user.Password,
	})
}
