package controllers

import (
	"net/http"

	"github.com/bubo-py/goonline/database"
	"github.com/bubo-py/goonline/models"
	"github.com/gin-gonic/gin"
)

// GetPokemon returns a json with pokemon with given id
func GetPokemon(c *gin.Context) {
	var pokemon models.Pokemon

	if err := database.Instance.Where("id = ?", c.Param("id")).First(&pokemon).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No profile with given ID"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"respond": pokemon})
}
