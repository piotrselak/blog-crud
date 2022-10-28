package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func HomeHandler(client *mongo.Collection) gin.HandlerFunc {
	handler := func(c *gin.Context) {
		c.String(http.StatusOK, "Front page") // placeholder for now
	}
	return gin.HandlerFunc(handler)
}
