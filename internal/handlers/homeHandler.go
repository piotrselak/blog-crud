package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piotrselak/blog-crud/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func HomeHandler(coll *mongo.Collection) gin.HandlerFunc {
	handler := func(c *gin.Context) {
		result := repository.GetAllPosts(coll)
		c.JSON(http.StatusOK, result)
	}
	return gin.HandlerFunc(handler)
}
