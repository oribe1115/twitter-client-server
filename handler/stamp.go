package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/oribe1115/twitter-client-server/model"
)

func CreateTableHandler(c echo.Context) error {
	err := model.CreateTable()

	if err != nil {
		return c.String(http.StatusInternalServerError, "faild to craete table\n")
	}
	return c.String(http.StatusCreated, "tasks table created!\n")
}

func AddNewStampHandler(c echo.Context) error {
	tweetid, _ := strconv.ParseInt(c.Param("tweetID"), 10, 64)
	stampid := c.Param("stampID")
	_, err := model.AddToStamp(tweetid, stampid)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to add")
	}

	stamplist, err := model.GetStampList(tweetid, stampid)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to get stamp list")
	}

	return c.JSON(http.StatusOK, stamplist)
}

func GetStampListHandler(c echo.Context) error {
	tweetid, _ := strconv.ParseInt(c.Param("tweetID"), 10, 64)
	stampid := c.Param("stampID")
	stamplist, err := model.GetStampList(tweetid, stampid)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to get stamp list")
	}

	return c.JSON(http.StatusOK, stamplist)
}

func DeleteToStampHandler(c echo.Context) error {
	tweetid, _ := strconv.ParseInt(c.Param("tweetID"), 10, 64)
	stampid := c.Param("stampID")
	err := model.DeleteToStamp(tweetid, stampid)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to delete stamp")
	}
	return c.String(http.StatusOK, "stamp is updated")
}
