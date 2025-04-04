package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tailnet_backend_go/models"
)

func RegisterNewUser(c *gin.Context) {
	var newUser models.UserRegister
	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Some field is missing"})
		return
	}
}
