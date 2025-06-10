package utils

import (
	"io"
	"os"
	"strings"
)

// ANSI Color Codes
const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorRed    = "\033[31m"
)

// GinColorWriter custom writer to colorize Gin's specific output lines
type GinColorWriter struct {
	Writer io.Writer
}

// Write method for our custom writer
func (w *GinColorWriter) Write(p []byte) (n int, err error) {
	s := string(p)
	if strings.Contains(s, "[GIN-debug] Listening and serving HTTP on") || strings.Contains(s, "Listening and serving HTTP on") {
		s = ColorGreen + s + ColorReset
	} else if strings.Contains(s, "[GIN-debug] [WARNING]") {
		s = ColorYellow + s + ColorReset
	} else if strings.Contains(s, "[GIN-debug] GET") || strings.Contains(s, "[GIN-debug] POST") ||
		strings.Contains(s, "[GIN-debug] PUT") || strings.Contains(s, "[GIN-debug] DELETE") { // Colorize routes
		s = ColorBlue + s + ColorReset
	}
	return w.Writer.Write([]byte(s))
}

// GetEnv is a helper to get environment variables with a default value
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}