package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/oribe1115/twitter-client-server/model"
)

func GetListsHandler(c echo.Context) error {
	lists, err := model.GetLists()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to get lists")
	}
	return c.JSON(http.StatusOK, lists)
}
