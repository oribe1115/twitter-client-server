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

func Retweet(api *anaconda.TwitterApi, tweetID int64) (StampTweet, error) {
	tweet, err := api.Retweet(tweetID, false)
	stampTweet := StampTweet{}

	if err != nil {
		return stampTweet, err
	}

	stampTweet.Tweet = tweet
	stampList, err := GetStampList(tweetID)

	if err != nil {
		return stampTweet, err
	}

	stampTweet.Stamp = stampList

	return stampTweet, nil
}

func Unretweet(api *anaconda.TwitterApi, tweetID int64) (StampTweet, error) {
	tweet, err := api.UnRetweet(tweetID, false)
	stampTweet := StampTweet{}

	if err != nil {
		return stampTweet, err
	}

	stampTweet.Tweet = tweet
	stampList, err := GetStampList(tweetID)

	if err != nil {
		return stampTweet, err
	}

	stampTweet.Stamp = stampList

	return stampTweet, nil
}

func Favorite(api *anaconda.TwitterApi, tweetID int64) (StampTweet, error) {
	stampTweet := StampTweet{}
	tweet, err := api.Favorite(tweetID)

	if err != nil {
		return stampTweet, err
	}

	stampTweet.Tweet = tweet
	stampList, err := GetStampList(tweetID)

	if err != nil {
		return stampTweet, err
	}

	stampTweet.Stamp = stampList

	return stampTweet, nil
}

func Unfavorite(api *anaconda.TwitterApi, tweetID int64) (StampTweet, error) {
	stampTweet := StampTweet{}
	tweet, err := api.Unfavorite(tweetID)

	if err != nil {
		return stampTweet, err
	}

	stampTweet.Tweet = tweet
	stampList, err := GetStampList(tweetID)

	if err != nil {
		return stampTweet, err
	}

	stampTweet.Stamp = stampList

	return stampTweet, nil
}
