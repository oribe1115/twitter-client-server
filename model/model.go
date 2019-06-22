package model

import (
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
	"github.com/joho/godotenv"
)

var (
	api          *anaconda.TwitterApi
	credentional *oauth.Credentials
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Twitter Apiのためのkeyなどをセットする
func GetTwitterAPI() {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	api = anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}
