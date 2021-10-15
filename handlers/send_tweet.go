package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joaomarcuslf/keyforge-name-of-the-day/clients"
	"github.com/joaomarcuslf/keyforge-name-of-the-day/usecases"
)

func SendTweet(c *gin.Context) {
	url := os.Getenv("KEYFORGE_URL")

	var client = usecases.NewKFRClient(url)

	kfr, err := client.GetRandom()

	if err != nil {
		panic(err)
	}

	twitterClient := clients.NewTwitterClient()
	twitterClient.Init()

	message := "The name of the day is " + kfr.Name + "! #Keyforge #KeyforgeNameOfTheDay"

	for _, house := range kfr.Links.Houses {
		message += " #" + house
	}

	twitterClient.SendTweet(message)

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Tweet sent",
		},
	)
}
