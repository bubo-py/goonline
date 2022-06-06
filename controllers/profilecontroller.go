package controllers

import (
	"net/http"

	"github.com/bubo-py/goonline/database"
	"github.com/bubo-py/goonline/models"
	"github.com/gin-gonic/gin"
)

// RegisterProfile created a new profile from provided data
func RegisterProfile(c *gin.Context) {
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Password hashed with HS256
	if err := profile.HashPassword(profile.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	record := database.Instance.Create(&profile)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"profileId": profile.ID,
		"profileLogin": profile.Login, "profileName": profile.Name,
		"profileSurname": profile.Surname, "profileAge": profile.Age})
}
