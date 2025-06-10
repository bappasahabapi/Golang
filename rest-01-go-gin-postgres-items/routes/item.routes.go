package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-postgres-local/controllers"
)

func RegisterItemRoutes(router *gin.Engine) {
	router.GET("/items", controllers.GetItems)
	router.GET("/items/:id", controllers.GetItem)
	router.POST("/items", controllers.CreateItem)
	router.PUT("/items/:id", controllers.UpdateItem)
	router.DELETE("/items/:id", controllers.DeleteItem)
}
