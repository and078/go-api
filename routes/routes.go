package routes

import (
	"btc_api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	deviceRoutes := r.Group("/devices")
	{
		deviceRoutes.GET("/", controllers.ShowDevices)
		deviceRoutes.GET("/:id", controllers.GetDeviceByIdController)
		deviceRoutes.POST("/", controllers.AddDevice)
                deviceRoutes.DELETE("/:id", controllers.DeleteDeviceByIdController)
	}

	fibonacciRoutes := r.Group("/fibonacci")
	{
		fibonacciRoutes.GET("/:number", controllers.FibonacciController)
	}
	return r
}
