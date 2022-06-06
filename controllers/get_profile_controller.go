package controllers

import (
	"net/http"

	"github.com/bubo-py/goonline/database"
	"github.com/bubo-py/goonline/models"
	"github.com/gin-gonic/gin"
)

// GetProfile returns a json with profile with given id
func GetProfile(c *gin.Context) {
	var profile models.Profile

	if err := database.Instance.Where("id = ?", c.Param("id")).First(&profile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No profile with given ID"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"respond": profile})
}
