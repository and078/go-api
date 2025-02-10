package controllers

import (
	"btc_api/initializers"
	"btc_api/models"

	"github.com/gin-gonic/gin"
)

func AddDevice (c *gin.Context) {
	var query struct {
		Name string
		Body string
	}

	c.Bind(&query)

	device := models.Device{Name: query.Name, Body: query.Body}

	result := initializers.DB.Create(&device)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error,
		})
		return
	}
	c.JSON(200, gin.H{
		"name": query.Name,
		"body": query.Body,
	})
}