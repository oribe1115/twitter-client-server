package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/oribe1115/twitter-client-server/model"
)

func TellMeHandler(c echo.Context) error {
	api := model.ApiFromContext(c)
	myDataToCheck, err := model.TellMe(api)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, myDataToCheck)
	}

	return c.JSON(http.StatusOK, myDataToCheck)
}

func TellOtherUserData(c echo.Context) error {
	api := model.ApiFromContext(c)

	userScreenName := c.Param("userScreenName")
	userData, err := model.TellOtherUserData(api, userScreenName)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to get user data")
	}

	return c.JSON(http.StatusOK, userData)
}
