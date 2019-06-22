package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/oribe1115/twitter-client-server/model"
)

func TellMeHandler(c echo.Context) error {
	myDataToCheck, err := model.TellMe()

	return c.JSON(http.StatusOK, myDataToCheck)
}
