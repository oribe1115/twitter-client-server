package model

import (
	"errors"
	"time"
)

//スタンプの構造体
type Stamp struct {
	StampID string 'json:"stamp_id"'   //twemojiの絵文字のid
	TweetID int64 'json:"tweet_id"'    //ツイートのid
	UserID int64 'json:"user_id"'		//スタンプを押した人のユーザーid
	UserScreenName string 'json:"user_screen_name"'    //スタンプを押した人のユーザー名(@..)
	Count int 'json:"count"'            //推されたスタンプの数
	CreatedAt string 'json:"created_at"'		//最初に押した日時
	UpdatedAt string 'json:"updated_at"'		//最後に押した日時
}

func createTable() error {
	err := db.createTable(&Stamp{}).Error
	if err !=nil{
		return err
	}

	return err
}

func AddToStamp(tweetId int64, stampID string) error {
	//既に押してあるか押してないかで分岐
	if(db.First(&stamp, "stamp_id=stampId", "tweet_id=tweetID") )


	//新しく追加されるスタンプレコードの作成
	newStamp := Stamp{}
	newStamp.StampID = stampID
	newStamp.TweetID = tweetId




	db.Table(stamps).Update("Count", ++)
	
}