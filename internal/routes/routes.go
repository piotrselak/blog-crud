package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/piotrselak/blog-crud/internal/db"
	"github.com/piotrselak/blog-crud/internal/handlers"
)

func RunRoutes() {
	client := db.InitializeMongoDB()
	postsCollection := client.Database("blog").Collection("posts")

	routes := gin.Default()
	routes.GET("/", handlers.HomeHandler(postsCollection))
	routes.Run()
}
