package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/oribe1115/twitter-client-server/model"
)

func TellMeHandler(c echo.Context) error {
	myDataToCheck, err := model.TellMe()

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, myDataToCheck)
	}

	return c.JSON(http.StatusOK, myDataToCheck)
}
