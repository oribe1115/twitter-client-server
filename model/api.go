package model

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/labstack/echo"
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

func ApiFromContext(c echo.Context) *anaconda.TwitterApi {
	token := c.Get("ACCESS_TOKEN").(string)
	secret := c.Get("ACCESS_TOKEN_SECRET").(string)

	api := anaconda.NewTwitterApiWithCredentials(token, secret, os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))

	return api
}
