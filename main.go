package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/oribe1115/twitter-client-server/handler"
	"github.com/oribe1115/twitter-client-server/model"
)

func main() {
	database := os.Getenv("DATABASE_URL")
	if database != "" {
		_, err := model.EstablishConnection()
		if err != nil {
			log.Fatal("Cannot Connect to Database: %s", err)
		}
	}

	enviroment := os.Getenv("ENVIROMENT")
	if enviroment == "" {
		model.LoadEnv()
	}

	token := os.Getenv("ACCESS_TOKEN")
	if token != "" {
		model.GetTwitterAPI()
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowCredentials: true,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HelloWorld")
	})

	e.GET("/user/me", handler.TellMeHandler)

	e.GET("/authorize", handler.GetRequestTokenHandler)
	e.GET("/authorize/callback", handler.GetAccessTokenHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	e.Start(":" + port)

}
