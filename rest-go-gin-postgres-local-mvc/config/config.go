package config

import "go-gin-postgres-local/utils" // Adjust module path if different

// DatabaseConfig holds database connection parameters
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// AppConfig holds all application configuration
type AppConfig struct {
	Database DatabaseConfig
	// Add other configurations here, e.g., ServerPort
}

// LoadConfig loads application configuration from environment variables or defaults
func LoadConfig() *AppConfig {
	return &AppConfig{
		Database: DatabaseConfig{
			Host:     utils.GetEnv("DB_HOST", "localhost"),
			Port:     utils.GetEnv("DB_PORT", "5432"),
			User:     utils.GetEnv("DB_USER", "postgres"),
			Password: utils.GetEnv("DB_PASSWORD", "postgres"),
			DBName:   utils.GetEnv("DB_NAME", "go-mvc"),
			SSLMode:  utils.GetEnv("DB_SSLMODE", "disable"),
		},
	}
}