package controllers

import (
	"blogAPI/database"
	"blogAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /tags
// Get all tags
func FindTags(c *gin.Context){
	var tags []models.Tag
	database.Db.Find(&tags)

	c.JSON(http.StatusOK, gin.H{"tags": tags})
}

// GET /tag