package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

func RetweetHandler(c echo.Context) error {
	api := model.ApiFromContext(c)

	tweetIDStr := c.Param("tweetID")
	tweetID, _ := strconv.ParseInt(tweetIDStr, 10, 64)

	stampTweet, err := model.Retweet(api, tweetID)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to unretweet")
	}

	return c.JSON(http.StatusOK, stampTweet)
}

func UnretweetHandler(c echo.Context) error {
	api := model.ApiFromContext(c)

	tweetIDStr := c.Param("tweetID")
	tweetID, _ := strconv.ParseInt(tweetIDStr, 10, 64)

	stampTweet, err := model.Unretweet(api, tweetID)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to retweet")
	}

	return c.JSON(http.StatusOK, stampTweet)
}
