package model

import (
	"github.com/ChimeraCoder/anaconda"
)

// 新規ツイート用
type NewTweet struct {
	Text string `json:"text,omitempty"`
}

// スタンプ情報付きツイート情報
type StampTweet struct {
	Tweet anaconda.Tweet
	Stamp
}

// 新規ツイートを投稿
func PostNewTweet(newTweet NewTweet) (StampTweet, error) {
	tweet, err := api.PostTweet(newTweet.Text, nil)
	stampTweet := StampTweet{}
	stampTweet.Tweet = tweet
	if err != nil {
		return stampTweet, err
	}

	return stampTweet, nil
}

func GetJustTweet(tweetID int64) (anaconda.Tweet, error) {
	tweet, err := api.GetTweet(tweetID, nil)
	if err != nil {
		return tweet, err
	}
	return tweet, nil
}
