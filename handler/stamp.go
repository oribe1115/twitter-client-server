package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/oribe1115/twitter-client-server/model"
)

// func CreateTableHandler(c echo.Context) error {
// 	err := model.CreateTable()

// 	if err != nil {
// 		return c.String(http.StatusInternalServerError, "faild to craete table\n")
// 	}
// 	return c.String(http.StatusCreated, "tasks table created!\n")
// }

// // スタンプを追加する
// func AddNewStampHandler(c echo.Context) error {
// 	tweetid, _ := strconv.ParseInt(c.Param("tweetID"), 10, 64)
// 	stampid := c.Param("stampID")
// 	stamp, err := model.AddToStamp(tweetid, stampid)
// 	if err != nil {
// 		fmt.Println(err)
// 		return c.String(http.StatusInternalServerError, "faild to add stamp")
// 	}
// 	tweet, err := model.GetJustTweet(tweetid)
// 	if err != nil {
// 		fmt.Println(err)
// 		return c.String(http.StatusInternalServerError, "faild get tweet")
// 	}
// 	tweetData := model.StampTweet{}
// 	tweetData.Tweet = tweet
// 	stampList := make([]model.Stamp, 0)
// 	stampList = append(stampList, stamp)
// 	tweetData.Stamp = stampList

// 	return c.JSON(http.StatusOK, tweetData)
// }

func GetStampListHandler(c echo.Context) error {
	tweetid, _ := strconv.ParseInt(c.Param("tweetID"), 10, 64)
	stamplist, err := model.GetStampList(tweetid)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to get stamp list")
	}

	return c.JSON(http.StatusOK, stamplist)
}

// func DeleteToStampHandler(c echo.Context) error {
// 	tweetid, _ := strconv.ParseInt(c.Param("tweetID"), 10, 64)
// 	stampid := c.Param("stampID")
// 	err := model.DeleteToStamp(tweetid, stampid)
// 	if err != nil {
// 		fmt.Println(err)
// 		return c.String(http.StatusInternalServerError, "faild to delete stamp")
// 	}
// 	return c.String(http.StatusOK, "stamp is updated")
// }
