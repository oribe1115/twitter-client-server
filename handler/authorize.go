package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
	"github.com/labstack/echo"

	"github.com/oribe1115/twitter-client-server/model"
)

var (
	credential   *oauth.Credentials
	apiInHandler *anaconda.TwitterApi
)

// 認証用urlをJSON形式で返す用
type AuthorizeURL struct {
	URL string `json:"url"`
}

// 認証用urlを返す
func GetRequestTokenHandler(c echo.Context) error {
	apiInHandler = anaconda.NewTwitterApiWithCredentials("", "", os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))

	url, tmpCred, err := apiInHandler.AuthorizationURL(os.Getenv("CALLBACK_URL"))

	fmt.Println(url)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to send authorizeing request")
	}
	credential = tmpCred

	authorizeURL := AuthorizeURL{}
	authorizeURL.URL = url

	fmt.Println("success to send authorizeing request")
	return c.JSON(http.StatusOK, authorizeURL)
}

// callback urlに帰ってきたものからtokenを取得
func GetAccessTokenHandler(c echo.Context) error {
	verifier := c.QueryParam("oauth_verifier")

	tmpCred, _, err := apiInHandler.GetCredentials(credential, verifier)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to get access token")
	}

	session, err := model.Store.Get(c.Request(), "session-key")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to get session")
	}

	session.Values["ACCESS_TOKEN"] = tmpCred.Token
	session.Values["ACCESS_TOKEN_SECRET"] = tmpCred.Secret

	err = session.Save(c.Request(), c.Response().Writer)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to save session")
	}

	fmt.Println("success to get access token")
	return c.String(http.StatusOK, "success to get access token")
}
