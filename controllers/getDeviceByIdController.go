package controllers

import (
	"btc_api/initializers"
	"btc_api/models"

	"github.com/gin-gonic/gin"
)

func GetDeviceByIdController(c *gin.Context) {
	id := c.Param("id")
	var device models.Device
	initializers.DB.First(&device, id)

	c.JSON(200, gin.H{
		"device": device,
	})
}
