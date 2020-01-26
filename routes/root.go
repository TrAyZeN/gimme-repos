package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to gimme-repos API, check routes field to see the available routes.",
		"routes": gin.H{
			"/reallyunknown": "Repositories with less than 20 stars.",
			"/unknown": "Repositories between 20 and 100 stars.",
			"/maybeknown": "Repositories between 100 and 1000 stars.",
			"/known": "Repositories with more than 1000 stars.",
		},
	})
}
