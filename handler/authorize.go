package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
	"github.com/labstack/echo"

	"github.com/oribe1115/practice-twitter/model"
)

var (
	credential   *oauth.Credentials
	apiInHandler *anaconda.TwitterApi
)

func GetRequestTokenHandler(c echo.Context) error {
	apiInHandler = anaconda.NewTwitterApiWithCredentials("", "", os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))

	url, tmpCred, err := apiInHandler.AuthorizationURL(os.Getenv("CALLBACK_URL"))

	fmt.Println(url)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to send authorizeing request")
	}
	credential = tmpCred

	fmt.Println("success to send authorizeing request")
	return c.String(http.StatusOK, url)
}

func GetAccessTokenHandler(c echo.Context) error {
	verifier := c.QueryParam("oauth_verifier")

	tmpCred, _, err := apiInHandler.GetCredentials(credential, verifier)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to get access token")
	}
	apiInHandlerWithToken := anaconda.NewTwitterApiWithCredentials(tmpCred.Token, tmpCred.Secret, os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	model.SetAPI(apiInHandlerWithToken)

	fmt.Println("success to get access token")
	return c.String(http.StatusOK, "success to get access token")
}
