package main

import (
	"btc_api/initializers"
	"btc_api/routes"
)

func init () {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main () {
	r := routes.SetupRouter()

	r.Run()
}