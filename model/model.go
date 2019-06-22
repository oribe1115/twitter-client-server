package model

import (
	"log"
	"os"
	"errors"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
)

var (
	api          *anaconda.TwitterApi
	credentional *oauth.Credentials
	db          *gorm.DB
	databaseURL string
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


// DBとの接続
func EstablishConnection() (*gorm.DB, error) {
	databaseURL = os.Getenv("DATABASE_URL")
	_db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		return nil, errors.New("faild to connect to DB")
	}
	db = _db

	return db, nil
}
