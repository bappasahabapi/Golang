package main

import (
	"log"
	"os"

	"go-gin-postgres-local/config"
	"go-gin-postgres-local/db"
	"go-gin-postgres-local/handlers"
	"go-gin-postgres-local/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // Import godotenv
)

func main() {
	// Load .env file at the very beginning
	err := godotenv.Load()
	if err != nil {
		// Log a warning if .env is not found, but don't make it fatal
		// as environment variables might be set directly in production.
		log.Println(utils.ColorYellow + "Warning: .env file not found or error loading it. Relying on preset environment variables." + utils.ColorReset)
	}

	// Load Application Configuration (this will now pick up from .env or system env)
	cfg := config.LoadConfig()

	// Connect to Database
	dbConn, dbErr := db.ConnectDB(&cfg.Database) // Renamed err to dbErr to avoid conflict
	if dbErr != nil {
		log.Fatalf(utils.ColorRed+"Failed to connect to database: %v"+utils.ColorReset, dbErr)
	}
	defer dbConn.Close()
	log.Println(utils.ColorGreen + "Successfully connected to local PostgreSQL!" + utils.ColorReset)

	// Initialize Database Schema
	if err := db.InitializeSchema(dbConn); err != nil { // Re-using err here is fine
		log.Fatalf(utils.ColorRed+"Error initializing database schema: %v"+utils.ColorReset, err)
	}

	// Setup Gin
	gin.DefaultWriter = &utils.GinColorWriter{Writer: os.Stdout}
	router := gin.Default()

	// Initialize Handlers
	itemHandler := handlers.NewItemHandler(dbConn)

	// Define API Routes
	itemRoutes := router.Group("/items")
	{
		itemRoutes.GET("", itemHandler.GetItems)
		itemRoutes.POST("", itemHandler.CreateItem)
		itemRoutes.GET("/:id", itemHandler.GetItem)
		itemRoutes.PUT("/:id", itemHandler.UpdateItem)
		itemRoutes.DELETE("/:id", itemHandler.DeleteItem)
	}

	// Start the server
	// Update config.go to include ServerPort if you want to manage it there too
	// For now, using utils.GetEnv directly here as before
	serverPort := utils.GetEnv("SERVER_PORT", "8080")
	if val := os.Getenv("SERVER_PORT_FROM_CONFIG"); val != "" { // Example if you had it in AppConfig
		// serverPort = cfg.ServerPort // Assuming you add ServerPort to AppConfig
	}

	log.Printf(utils.ColorGreen+"Starting server on :%s"+utils.ColorReset, serverPort)
	if err := router.Run(":" + serverPort); err != nil {
		log.Fatalf(utils.ColorRed+"Failed to run server: %v"+utils.ColorReset, err)
	}
}