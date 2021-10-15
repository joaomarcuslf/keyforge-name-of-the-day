package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaomarcuslf/keyforge-name-of-the-day/usecases"
)

func SendTweet(c *gin.Context) {
	kfr, err := usecases.SendTweet()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message":   "Tweet sent",
			"deck_name": kfr.Name,
		},
	)
}
