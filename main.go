package main

import (
	"log"

	"github.com/joho/godotenv"
	"linkbulb.io/links-service/models"
	"linkbulb.io/links-service/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.Init()
	routes.Init()
}
