package main

import (
	"btc_api/controllers"
	"btc_api/initializers"

	"github.com/gin-gonic/gin"
)

func init () {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main () {
	r := gin.Default()

	r.POST("/devices", controllers.AddDevice)
	r.GET("/devices", controllers.ShowDevices)
	r.GET("/device/:id", controllers.GetDeviceByIdController)

	r.Run()
}