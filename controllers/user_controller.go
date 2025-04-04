package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tailnet_backend_go/models"
	"tailnet_backend_go/services"
)

func RegisterUser(c *gin.Context) {

	var newUser models.UserRegister

	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing some fields"})
		return
	}

	err := services.RegisterNewUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "User registered successfully"})
}
