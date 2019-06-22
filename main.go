package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/oribe1115/twitter-client-server/model"

)

func main() {
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	e.Start(":" + port)

}
