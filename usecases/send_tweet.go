package usecases

import (
	"os"
	"strings"

	"github.com/joaomarcuslf/keyforge-name-of-the-day/clients"
)

func SendTweet() (KFRDeck, error) {
	url := os.Getenv("KEYFORGE_URL")

	var client = NewKFRClient(url)

	kfr, err := client.GetRandom()

	if err != nil {
		return KFRDeck{}, err
	}

	twitterClient := clients.NewTwitterClient()
	twitterClient.Init()

	message := "The name of the day is " + kfr.Name + "! #Keyforge #KeyforgeNameOfTheDay"

	for _, house := range kfr.Links.Houses {
		message += " #" + strings.Replace(house, " ", "_", -1)
	}

	tweet, response, err := twitterClient.SendTweet(message)

	if err != nil {
		return KFRDeck{}, err
	}

	if response.StatusCode != 200 {
		return KFRDeck{}, err
	}

	kfr.TweetUrl = "https://twitter.com/KFDeckEveryday/status/" + tweet.IDStr

	return kfr, nil
}
