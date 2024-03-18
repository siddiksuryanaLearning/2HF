package main

import (
	"2hf/config"
	"2hf/docs"
	"2hf/routes"
	"2hf/utils"

	"log"

	"github.com/joho/godotenv"
)

func main() {

	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Halal Foods."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.Getenv("SWAGGER_HOST", "18.219.232.22:8080")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
	r.Run()
}
