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
	Stamp []Stamp
}

// 新規ツイートを投稿
// あとでstampのデフォルトデータを追加
func PostNewTweet(api *anaconda.TwitterApi, newTweet NewTweet) (StampTweet, error) {
	tweet, err := api.PostTweet(newTweet.Text, nil)

	stampTweet := StampTweet{
		Tweet: tweet,
	}

	if err != nil {
		return stampTweet, err
	}

	return stampTweet, nil
}

func GetJustTweet(api *anaconda.TwitterApi, tweetID int64) (anaconda.Tweet, error) {
	tweet, err := api.GetTweet(tweetID, nil)

	if err != nil {
		return tweet, err
	}

	return tweet, nil
}
