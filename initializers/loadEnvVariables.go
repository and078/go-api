package initializers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables () {
	fmt.Println("LoadEnvVariables")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}