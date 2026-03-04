package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/mallacharmi/polyglot-export-engine/internal/database"
	"github.com/mallacharmi/polyglot-export-engine/internal/handlers"
	"github.com/mallacharmi/polyglot-export-engine/internal/services"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Connect to database
	database.Connect()

	router := gin.Default()

	exportService := services.NewExportService()
	exportHandler := handlers.NewExportHandler(exportService)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	router.POST("/exports", exportHandler.CreateExport)
	router.GET("/exports/:id", exportHandler.GetExport)

	log.Println("Server running on port", port)
	router.Run(":" + port)
}