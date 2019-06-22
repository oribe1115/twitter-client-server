package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/oribe1115/twitter-client-server/model"
)

func GetHomeTimelineHandler(c echo.Context) error {
	countStr := c.QueryParam("count")
	if countStr == "" {
		countStr = "20"
	}
	count, _ := strconv.Atoi(countStr)
	homeTimeline, err := model.GetHomeTimeline(count)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to get home timeline")
	}

	return c.JSON(http.StatusOK, homeTimeline)
}

func GetUserTimelineHandler(c echo.Context) error {
	countStr := c.QueryParam("count")
	if countStr == "" {
		countStr = "20"
	}
	count, _ := strconv.Atoi(countStr)

	userID := c.Param("userID")

	stampTweetList, err := model.GetUserTimeline(userID, count)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to get user timeline")
	}

	return c.JSON(http.StatusOK, stampTweetList)
}
