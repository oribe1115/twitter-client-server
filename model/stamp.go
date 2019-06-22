package model

import (
	"errors"
	"time"
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

func createTable() error {
	err := db.CreateTable(&Stamp{}).Error
	if err != nil {
		return err
	}

	return err
}

func AddToStamp(tweetID int64, stampID string) error {
	userID, _ := TellMyUserId()
	userScreenName, _ := TellMyScreenName()
	var checkCount int
	var checkStamp Stamp
	db.Where("stamp_id = ? AND tweet_id = ? AND user_id = ? And user_screen_name = ?", stampID, tweetID, userID, userScreenName).First(&checkStamp).Count(&checkCount)

	//既に押してあるか押してないかで分岐
	if checkCount == 0 {
		//新しく追加されるスタンプレコードの作成
		newStamp := Stamp{}
		newStamp.StampID = stampID
		newStamp.TweetID = tweetID
		newStamp.UserID, _ = TellMyUserId()
		newStamp.UserScreenName, _ = TellMyScreenName()
		newStamp.Count = 1
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		newStamp.CreatedAt = time.Now().In(jst)
		newStamp.UpdatedAt = time.Now().In(jst)

		//DBにデータを挿入
		err := db.Table("stamps").Create(&newStamp).Error
		if err != nil {
			return errors.New("faild to add task")
		}

		return nil

	} else {
		//countを1増やす
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		checkStamp.UpdatedAt = time.Now().In(jst)
		checkStamp.Count++
		db.Save(&checkStamp)

	}

}

func DeleteToStamp(tweetId int64, stampID string) error {

}
