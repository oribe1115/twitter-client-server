package handler

import (
	"net/http"

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
