package main

import (
	"os"

	db "github.com/Ademayowa/learn-d-compose/internal/database"
	"github.com/Ademayowa/learn-d-compose/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		println("No .env file found, using default environment variables")
	}

	db.InitDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := gin.Default()
	handlers.RegisterRoutes(server)

	server.Run(":" + port)
}
