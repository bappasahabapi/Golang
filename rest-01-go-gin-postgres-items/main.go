package main

import (
	"log"
	"go-gin-postgres-local/config"
	"go-gin-postgres-local/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	router := gin.Default()

	routes.RegisterItemRoutes(router)

	log.Println("🚀 Server running at http://localhost:8080")
	router.Run(":8080")
}
