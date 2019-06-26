package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/oribe1115/twitter-client-server/model"
)

func CheckAuthorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, err := model.Store.Get(c.Request(), "session-key")
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "something wrong in getting session")
		}

		if session.Values["ACCESS_TOKEN"] == nil || session.Values["ACCESS_TOKEN_SECRET"] == nil {
			return c.String(http.StatusForbidden, "please authorize with twitter")
		}

		c.Set("ACCESS_TOKEN", session.Values["ACCESS_TOKEN"].(string))
		c.Set("ACCESS_TOKEN_SECRET", session.Values["ACCESS_TOKEN_SECRET"].(string))

		return next(c)
	}
}
