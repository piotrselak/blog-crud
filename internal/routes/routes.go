package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/piotrselak/blog-crud/internal/handlers"
)

func RunRoutes() {
	routes := gin.Default()
	routes.GET("/", handlers.HomeHandler)
	routes.Run()
}
