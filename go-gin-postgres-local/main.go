package main

import (
	// "bytes" // To capture and modify Gin's output
	"database/sql"
	"fmt"
	"io" // For io.Writer
	"log"
	"net/http"
	"os"
	"strings" // For string manipulation

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// ANSI Color Codes
const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m" // For Gin's debug warnings
	ColorBlue   = "\033[34m" // For Gin's route info
	ColorRed    = "\033[31m" // For errors
)

var db *sql.DB

type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func initializeSchema(db *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS items (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		quantity INTEGER
	);`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating items table: %w", err)
	}
	log.Println(ColorGreen + "Items table checked/created successfully." + ColorReset) // Green color
	return nil
}

// Custom writer to colorize Gin's specific output lines
type ginColorWriter struct {
	writer io.Writer
}

// Write method for our custom writer
func (w *ginColorWriter) Write(p []byte) (n int, err error) {
	s := string(p)
	if strings.Contains(s, "[GIN-debug] Listening and serving HTTP on") || strings.Contains(s, "Listening and serving HTTP on") {
		s = ColorGreen + s + ColorReset
	} else if strings.Contains(s, "[GIN-debug] [WARNING]") {
		s = ColorYellow + s + ColorReset
	} else if strings.Contains(s, "[GIN-debug] GET") || strings.Contains(s, "[GIN-debug] POST") { // Colorize routes
		s = ColorBlue + s + ColorReset
	}
	// You can add more rules here for other Gin messages if desired

	return w.writer.Write([]byte(s))
}

func main() {
	// --- Database Connection Details ---
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "foodie")
	sslMode := getEnv("DB_SSLMODE", "disable")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf(ColorRed+"Error opening database: %v"+ColorReset, err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf(ColorRed+"Error connecting to database: %v. Check connection string and if PostgreSQL is running."+ColorReset, err)
	}
	// Apply green color to this log message
	log.Println(ColorGreen + "Successfully connected to local PostgreSQL!" + ColorReset)

	if err := initializeSchema(db); err != nil {
		log.Fatalf(ColorRed+"Error initializing database schema: %v"+ColorReset, err)
	}

	// --- Gin Setup with Colored Output ---
	// To colorize Gin's "Listening and serving..." message, we need to set its output writer.
	// gin.DefaultWriter = os.Stdout by default.
	// gin.DefaultErrorWriter = os.Stderr by default.

	// Create our custom colorizing writer targeting os.Stdout
	// This will affect Gin's standard logs, including the "Listening..." message
	coloredGinDefaultWriter := &ginColorWriter{writer: os.Stdout}
	gin.DefaultWriter = coloredGinDefaultWriter
	// You could also set gin.DefaultErrorWriter if you want to colorize Gin's error output

	router := gin.Default() // This will now use our custom writer for its default logger

	router.GET("/items", getItems)
	router.GET("/items/:id", getItem)
	router.POST("/items", createItem)

	// Apply green color to this log message
	// Note: Gin's "Listening and serving..." message will be colored by the custom writer.
	// This log.Println is for your custom "Starting server..." message.
	log.Println(ColorGreen + "Starting server on :8080" + ColorReset)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf(ColorRed+"Failed to run server: %v"+ColorReset, err)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// --- Gin Handlers (Modified slightly to use colored logging for errors) ---

func getItems(c *gin.Context) {
	rows, err := db.Query("SELECT id, name, quantity FROM items")
	if err != nil {
		log.Printf(ColorRed+"Error querying items: %v"+ColorReset, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve items"})
		return
	}
	defer rows.Close()

	items := []Item{}
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Quantity); err != nil {
			log.Printf(ColorRed+"Error scanning item row: %v"+ColorReset, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process item data"})
			return
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf(ColorRed+"Error after iterating rows: %v"+ColorReset, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing item results"})
		return
	}

	c.JSON(http.StatusOK, items)
}

func getItem(c *gin.Context) {
	id := c.Param("id")
	var item Item

	row := db.QueryRow("SELECT id, name, quantity FROM items WHERE id = $1", id)
	err := row.Scan(&item.ID, &item.Name, &item.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		} else {
			log.Printf(ColorRed+"Error querying single item: %v"+ColorReset, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve item"})
		}
		return
	}
	c.JSON(http.StatusOK, item)
}

func createItem(c *gin.Context) {
	var newItem Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newItem.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item name cannot be empty"})
		return
	}

	err := db.QueryRow(
		"INSERT INTO items (name, quantity) VALUES ($1, $2) RETURNING id",
		newItem.Name, newItem.Quantity,
	).Scan(&newItem.ID)

	if err != nil {
		log.Printf(ColorRed+"Error creating item: %v"+ColorReset, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	c.JSON(http.StatusCreated, newItem)
}