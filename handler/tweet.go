package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/oribe1115/twitter-client-server/model"
)

// 新規ツイートを投稿
func NewTweetPostHandler(c echo.Context) error {
	api := model.ApiFromContext(c)

	newTweet := model.NewTweet{}
	c.Bind(&newTweet)

	stampTweet, err := model.PostNewTweet(api, newTweet)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to tweet")
	}

	return c.JSON(http.StatusInternalServerError, stampTweet)
}
