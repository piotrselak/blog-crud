package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piotrselak/blog-crud/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func PostHandler(coll *mongo.Collection) gin.HandlerFunc {
	handler := func(c *gin.Context) {
		title := c.Query("title") // done for fast testing
		text := c.Query("text")
		result := repository.AddNewPost(coll, title, text)
		if result != nil {
			c.String(http.StatusBadRequest, "Bad request - validation failed")
		} else {
			c.String(http.StatusOK, "Post added")
		}

	}
	return gin.HandlerFunc(handler)
}
