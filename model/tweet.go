package model

import (
	"github.com/ChimeraCoder/anaconda"
)

// 新規ツイート用
type NewTweet struct {
	Text string `json:"text,omitempty"`
}

type StampTweet struct {
	Tweet anaconda.Tweet
	Stamp
}

func PostNewTweet(newTweet NewTweet) (StampTweet, error) {
	tweet, err := api.PostTweet(newTweet.Text, nil)
	stampTweet := StampTweet{}
	stampTweet.Tweet = tweet
	if err != nil {
		return stampTweet, err
	}

	return stampTweet, nil
}
