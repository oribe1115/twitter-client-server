package model

import (
	"time"

	"github.com/ChimeraCoder/anaconda"
)

//スタンプの構造体
type Stamp struct {
	StampID        string    `json:"stamp_id"`         //twemojiの絵文字のid
	TweetID        int64     `json:"tweet_id"`         //ツイートのid
	UserID         int64     `json:"user_id"`          //スタンプを押した人のユーザーid
	UserScreenName string    `json:"user_screen_name"` //スタンプを押した人のユーザー名(@..)
	Count          int       `json:"count"`            //推されたスタンプの数
	CreatedAt      time.Time `json:"created_at"`       //最初に押した日時
	UpdatedAt      time.Time `json:"updated_at"`       //最後に押した日時
}

func CreateTable() error {
	err := db.CreateTable(&Stamp{}).Error
	if err != nil {
		return err
	}

	return err
}

//スタンプを追加する関数
func AddStamp(api *anaconda.TwitterApi, tweetID int64, stampID string) (Stamp, error) {
	var checkCount int

	user, _ := api.GetSelf(nil)
	userID := user.Id

	checkStamp := Stamp{}
	db.Where("stamp_id = ? AND tweet_id = ? AND user_id = ?", stampID, tweetID, userID).First(&checkStamp).Count(&checkCount)

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	if checkCount == 0 {
		newStamp := Stamp{
			StampID:        stampID,
			TweetID:        tweetID,
			UserID:         userID,
			UserScreenName: user.ScreenName,
			Count:          1,
			CreatedAt:      time.Now().In(jst),
			UpdatedAt:      time.Now().In(jst),
		}

		err := db.Table("stamps").Create(&newStamp).Error
		if err != nil {
			return newStamp, err
		}

		return newStamp, nil

	}

	checkStamp.UpdatedAt = time.Now().In(jst)
	checkStamp.Count++
	db.Save(&checkStamp)

	return checkStamp, nil
}

//スタンプを削除する関数
func DeleteStamp(api *anaconda.TwitterApi, tweetID int64, stampID string) error {
	user, _ := api.GetSelf(nil)
	userID := user.Id
	stamp := Stamp{}

	db.Where("stamp_id = ? AND tweet_id = ? AND user_id = ?", stampID, tweetID, userID).First(&stamp)
	err := db.Delete(&stamp).Error

	if err != nil {
		return err
	}

	return nil
}

//スタンプのリストを取得する関数
func GetStampList(tweetID int64) ([]Stamp, error) {
	stampList := []Stamp{}
	err := db.Where("tweet_id = ?", tweetID).Select("*").Find(&stampList).Error

	if err != nil {
		return nil, err
	}

	return stampList, nil
}
