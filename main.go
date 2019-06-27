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
	enviroment := os.Getenv("ENVIROMENT")
	if enviroment == "" {
		model.LoadEnv()
	}

	database := os.Getenv("DATABASE_URL")
	if database != "" {
		_, err := model.EstablishConnection()
		if err != nil {
			log.Fatal("Cannot Connect to Database: %s", err)
		}
	}

	// token := os.Getenv("ACCESS_TOKEN")
	// if token != "" {
	// 	model.GetTwitterAPI()
	// }

	_, err := model.StoreForSession()
	if err != nil {
		log.Fatal("faild at session: %s", err)
	}

	e := echo.New()
	e.Use(middleware.CORS())
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins:     []string{"http://localhost:8080"},
	// 	AllowCredentials: true,
	// }))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HelloWorld")
	})
	// e.GET("/create/table", handler.CreateTableHandler)

	e.GET("/authorize", handler.GetRequestTokenHandler)
	e.GET("/authorize/callback", handler.GetAccessTokenHandler)

	withTwitter := e.Group("")
	withTwitter.Use(handler.CheckAuthorize)

	withTwitter.GET("/user/me", handler.TellMeHandler)
	withTwitter.GET("user/:userScreenName", handler.TellOtherUserData)

	withTwitter.POST("/tweet/:tweetID/stamps/:stampID", handler.AddNewStampHandler)
	withTwitter.GET("/tweet/:tweetID/stamps", handler.GetStampListHandler)
	withTwitter.POST("/tweet/:tweetID/stamps/:stampID/destroy", handler.DeleteToStampHandler)

	withTwitter.POST("tweet/new", handler.NewTweetPostHandler)
	withTwitter.POST("/tweet/:tweetID/retweet", handler.RetweetHandler)
	withTwitter.POST("/tweet/:tweetID/unretweet", handler.UnretweetHandler)

	withTwitter.GET("/statuses/home_timeline", handler.GetHomeTimelineHandler)
	withTwitter.GET("/lists/list", handler.GetListsHandler)
	withTwitter.GET("/lists/:listID/statuses", handler.GetListStatusesHandler)
	withTwitter.GET("/statuses/:userID", handler.GetUserTimelineHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	e.Start(":" + port)

}
