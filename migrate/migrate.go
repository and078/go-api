package main

import (
	"btc_api/initializers"
	"btc_api/models"
	"fmt"
)

func init () {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main () {
	if initializers.DB != nil {
		initializers.DB.AutoMigrate(&models.Device{})
		return
	}
	fmt.Println("Db is nil")
}