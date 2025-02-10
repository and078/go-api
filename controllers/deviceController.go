package controllers

import (
	"btc_api/initializers"
	"btc_api/models"

	"github.com/gin-gonic/gin"
) 

func ShowDevices (c *gin.Context) {
	var devices []models.Device
	var result = initializers.DB.Find(&devices)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, devices)
}