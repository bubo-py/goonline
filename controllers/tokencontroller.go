package controllers

import (
	"net/http"

	"github.com/bubo-py/goonline/auth"
	"github.com/bubo-py/goonline/database"
	"github.com/bubo-py/goonline/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// GenerateToken generates JWT
func GenerateToken(c *gin.Context) {
	var request TokenRequest
	var profile models.Profile

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	record := database.Instance.Where("login = ?", request.Login).First(&profile)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}

	credentialErr := profile.CheckPassword(request.Password)
	if credentialErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		c.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(profile.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Logout logs a profile out by deleting current cookie session and terminating the JWT
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}

	session.Delete(tokenString)

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
