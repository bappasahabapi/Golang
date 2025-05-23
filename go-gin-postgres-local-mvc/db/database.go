package db

import (
	"database/sql"
	"fmt"
	"log"

	"go-gin-postgres-local/config" // Adjust module path
	"go-gin-postgres-local/utils"  // Adjust module path

	_ "github.com/lib/pq" // PostgreSQL driver
)

// ConnectDB establishes a connection to the PostgreSQL database
func ConnectDB(cfg *config.DatabaseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close() // Close the connection if ping fails
		return nil, fmt.Errorf("error connecting to database: %w. Check connection string and if PostgreSQL is running", err)
	}
	return db, nil
}

// InitializeSchema creates necessary tables if they don't exist
func InitializeSchema(db *sql.DB) error {
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
	log.Println(utils.ColorGreen + "Items table checked/created successfully." + utils.ColorReset)
	return nil
}