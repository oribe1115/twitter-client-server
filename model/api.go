package model

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Twitter Apiのためのkeyなどをセットする
func GetTwitterAPI() {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	api = anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}

func SetAPI(apiInHandler *anaconda.TwitterApi) {
	api = apiInHandler
}
