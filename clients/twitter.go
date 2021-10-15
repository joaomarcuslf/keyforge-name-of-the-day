package clients

// import (
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/dghubble/go-twitter/twitter"
// 	"github.com/dghubble/oauth1"
// )

// type TwitterClient struct {
// 	client *twitter.Client
// 	creds  Credentials
// }

// type Credentials struct {
// 	ConsumerKey       string
// 	ConsumerSecret    string
// 	AccessToken       string
// 	AccessTokenSecret string
// }

// func NewTwitterClient() *TwitterClient {
// 	return &TwitterClient{
// 		creds: Credentials{
// 			AccessToken:       os.Getenv("ACCESS_TOKEN"),
// 			AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
// 			ConsumerKey:       os.Getenv("CONSUMER_KEY"),
// 			ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
// 		},
// 	}
// }

// func (t *TwitterClient) Init() {
// 	client, err := getClient(&t.creds)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	t.client = client
// }

// func (t *TwitterClient) SendTweet(message string) (*twitter.Tweet, *http.Response, error) {
// 	tweet, resp, err := t.client.Statuses.Update(message, nil)

// 	return tweet, resp, err
// }

// func getClient(creds *Credentials) (*twitter.Client, error) {
// 	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)

// 	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

// 	httpClient := config.Client(oauth1.NoContext, token)
// 	client := twitter.NewClient(httpClient)

// 	verifyParams := &twitter.AccountVerifyParams{
// 		SkipStatus:   twitter.Bool(true),
// 		IncludeEmail: twitter.Bool(true),
// 	}

// 	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
// 	if err != nil {
// 		return nil, err
// 	}

// 	log.Printf("User's ACCOUNT:\n%+v\n", user)
// 	return client, nil
// }
