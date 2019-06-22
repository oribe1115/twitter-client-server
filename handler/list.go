package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

func GetListStatusesHandler(c echo.Context) error {
	listID, _ := strconv.ParseInt(c.Param("listID"), 10, 64)
	countStr := c.QueryParam("count")
	if countStr == "" {
		countStr = "20"
	}
	count, _ := strconv.Atoi(countStr)

	listStatuses, err := model.GetListStatuses(listID, count)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to get list statuses")
	}

	return c.JSON(http.StatusOK, listStatuses)
}
