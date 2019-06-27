package model

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var (
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
		return nil, err
	}
	db = _db

	return db, nil
}
