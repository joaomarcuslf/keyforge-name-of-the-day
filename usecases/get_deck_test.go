package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var client = NewKFRClient("")

func init() {}

func TestGet01(t *testing.T) {
	response, err := client.Get(1, 9, "")

	assert.True(t, err == nil)
	assert.True(t, len(response.Data) == 9)
	assert.True(t, response.Count > 0)
}

func TestGet02(t *testing.T) {
	_, err := client.Get(0, 1, "")

	assert.True(t, err != nil)
	assert.Error(t, err)
}

func TestGet03(t *testing.T) {
	_, err := client.Get(1, 0, "")

	assert.True(t, err != nil)
	assert.Error(t, err)
}

func TestGetDecks(t *testing.T) {
	decks, err := client.GetDecks(1, 5)

	assert.True(t, err == nil)
	assert.True(t, len(decks) == 5)
}

func TestGetFirstDeck(t *testing.T) {
	deck, err := client.GetFirstDeck()

	assert.True(t, err == nil)

	assert.True(t, deck.Expansion > 0)
	assert.True(t, deck.Name != "")
}

func TestGetRandom(t *testing.T) {
	deck, err := client.GetRandom()

	assert.True(t, err == nil)

	assert.True(t, deck.Expansion > 0)
	assert.True(t, deck.Name != "")
}
