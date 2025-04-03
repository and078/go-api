package controllers

import (
	"btc_api/initializers"
	"btc_api/models"

	"github.com/gin-gonic/gin"
)

func DeleteDeviceByIdController(c *gin.Context){
	id := c.Param("id")
	var device models.Device
	initializers.DB.Delete(&device, id)

	c.JSON(200, gin.H{
		"deleted device": id,
	})
}
