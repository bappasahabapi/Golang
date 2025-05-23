package main

import (
	"log"
	"os"

	"go-gin-postgres-local/config"   // Adjust if your module path is different
	"go-gin-postgres-local/db"       // Adjust
	"go-gin-postgres-local/handlers" // Adjust
	"go-gin-postgres-local/utils"    // Adjust

	"github.com/gin-gonic/gin"
)

func main() {
	// Load Application Configuration
	cfg := config.LoadConfig()

	// Connect to Database
	dbConn, err := db.ConnectDB(&cfg.Database)
	if err != nil {
		log.Fatalf(utils.ColorRed+"Failed to connect to database: %v"+utils.ColorReset, err)
	}
	defer dbConn.Close()
	log.Println(utils.ColorGreen + "Successfully connected to local PostgreSQL!" + utils.ColorReset)

	// Initialize Database Schema
	if err := db.InitializeSchema(dbConn); err != nil {
		log.Fatalf(utils.ColorRed+"Error initializing database schema: %v"+utils.ColorReset, err)
	}

	// Setup Gin
	// Set Gin's default writer to our custom colorizing writer
	gin.DefaultWriter = &utils.GinColorWriter{Writer: os.Stdout}
	// gin.DefaultErrorWriter = &utils.GinColorWriter{Writer: os.Stderr} // Optional: colorize Gin's error output too

	router := gin.Default() // This will now use our custom writer

	// Initialize Handlers
	itemHandler := handlers.NewItemHandler(dbConn)

	// Define API Routes
	itemRoutes := router.Group("/items")
	{
		itemRoutes.GET("", itemHandler.GetItems)
		itemRoutes.GET("/:id", itemHandler.GetItem)
		itemRoutes.POST("", itemHandler.CreateItem)
		// Add PUT and DELETE handlers here similarly
	}

	// Start the server
	serverPort := utils.GetEnv("SERVER_PORT", "8080")
	log.Printf(utils.ColorGreen+"Starting server on :%s"+utils.ColorReset, serverPort)
	if err := router.Run(":" + serverPort); err != nil {
		log.Fatalf(utils.ColorRed+"Failed to run server: %v"+utils.ColorReset, err)
	}
}