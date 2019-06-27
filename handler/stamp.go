package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/oribe1115/twitter-client-server/model"
)

// 新規のテーブルstampsを作成する
func CreateTableHandler(c echo.Context) error {
	err := model.CreateTable()

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to craete table\n")
	}

	return c.String(http.StatusCreated, "tasks table created!\n")
}

// スタンプを追加する
func AddNewStampHandler(c echo.Context) error {
	api := model.ApiFromContext(c)

	tweetIDStr := c.Param("tweetID")
	tweetID, _ := strconv.ParseInt(tweetIDStr, 10, 64)
	stampID := c.Param("stampID")
	stamp, err := model.AddStamp(api, tweetID, stampID)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to add stamp")
	}

	tweet, err := model.GetJustTweet(api, tweetID)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild get tweet")
	}

	stampTweet := model.StampTweet{}
	stampTweet.Tweet = tweet
	stampList := make([]model.Stamp, 0)
	stampList = append(stampList, stamp)
	stampTweet.Stamp = stampList

	return c.JSON(http.StatusOK, stampTweet)
}

func GetStampListHandler(c echo.Context) error {
	tweetIDStr := c.Param("tweetID")
	tweetID, _ := strconv.ParseInt(tweetIDStr, 10, 64)
	stampList, err := model.GetStampList(tweetID)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to get stamp list")
	}

	return c.JSON(http.StatusOK, stampList)
}

func DeleteToStampHandler(c echo.Context) error {
	api := model.ApiFromContext(c)

	tweetIDStr := c.Param("tweetID")
	tweetID, _ := strconv.ParseInt(tweetIDStr, 10, 64)
	stampID := c.Param("stampID")
	err := model.DeleteStamp(api, tweetID, stampID)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to delete stamp")
	}

	return c.String(http.StatusOK, "stamp is deleted")
}
