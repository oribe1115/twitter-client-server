package model

import (
	"errors"
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var (
	api         *anaconda.TwitterApi
	db          *gorm.DB
	databaseURL string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// DBとの接続
func EstablishConnection() (*gorm.DB, error) {
	databaseURL = os.Getenv("DATABASE_URL")
	_db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		return nil, errors.New("faild to connect to DB")
	}
	db = _db

	return db, nil
}
