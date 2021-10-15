package usecases

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var KF_URL = "https://www.keyforgegame.com/"

type KFRDeck struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Expansion  int    `json:"expansion"`
	PowerLevel int    `json:"power_level"`
	Chains     int    `json:"chains"`

	Wins   int `json:"wins"`
	Losses int `json:"losses"`

	CasualWins   int `json:"casual_wins"`
	CasualLosses int `json:"casual_losses"`

	Cards           []string `json:"cards"`
	Notes           []string `json:"notes"`
	IsMyDeck        bool     `json:"is_my_deck"`
	IsMyFavorite    bool     `json:"is_my_favorite"`
	IsOnMyWatchlist bool     `json:"is_on_my_watchlist"`

	SetEraCards map[string]interface{} `json:"set_era_cards"`
	ShardsBonus interface{}            `json:"shards_bonus"`

	Links KFRLinks `json:"_links"`
}

type KFRLinks struct {
	Houses []string `json:"houses"`
}

type KFResponse struct {
	Data   []KFRDeck              `json:"data"`
	Linked map[string]interface{} `json:"_linked"`
	Count  int                    `json:"count"`
}

type KFRClient struct {
	URL string
}

func NewKFRClient(url string) *KFRClient {
	if url == "" {
		url = KF_URL
	}

	return &KFRClient{
		URL: url,
	}
}

func (c *KFRClient) Get(page, size int, ordering string) (KFResponse, error) {
	if page < 1 {
		return KFResponse{}, fmt.Errorf("page must be greater than 0")
	}

	if size < 1 {
		return KFResponse{}, fmt.Errorf("size must be greater than 0")
	}

	if ordering == "" {
		ordering = "-date"
	}

	res, err := http.Get(
		fmt.Sprintf(
			"%s/api/decks/?page=%d&page_size=%d&ordering=%s",
			c.URL,
			page,
			size,
			ordering,
		),
	)

	if err != nil {
		return KFResponse{}, err
	}

	var kfr KFResponse

	err = json.NewDecoder(res.Body).Decode(&kfr)

	if err != nil {
		return KFResponse{}, err
	}

	return kfr, nil
}

func (c *KFRClient) GetDecks(page, size int) ([]KFRDeck, error) {
	kfr, err := c.Get(page, size, "-date")

	if err != nil {
		return []KFRDeck{}, err
	}

	return kfr.Data, nil
}

func (c *KFRClient) GetFirstDeck() (KFRDeck, error) {
	kfr, err := c.Get(1, 1, "-date")

	if err != nil {
		return KFRDeck{}, err
	}

	return kfr.Data[0], nil
}

func (c *KFRClient) GetRandom() (KFRDeck, error) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 35000

	page := rand.Intn(max-min+1) + min

	min = 5
	max = 20

	size := rand.Intn(max-min+1) + min

	kfr, err := c.Get(page, size, "-date")

	if err != nil {
		return KFRDeck{}, err
	}

	min = 0
	max = len(kfr.Data) - 1

	random := rand.Intn(max-min+1) + min

	return kfr.Data[random], nil
}
