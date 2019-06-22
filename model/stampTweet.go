package model

import (
	"github.com/ChimeraCoder/anaconda"
)

// []Tweetを[]StampTweetにする
func MakeStampTweet(tweetList []anaconda.Tweet) ([]StampTweet, error) {
	stampTweetList := make([]StampTweet, 0)

	for _, tweet := range tweetList {
		stampList, err := GetStampList(tweet.Id)
		if err != nil {
			return nil, err
		}
		stampTweet := StampTweet{}
		stampTweet.Tweet = tweet
		stampTweet.Stamp = stampList
		stampTweetList = append(stampTweetList, stampTweet)

	}

	return stampTweetList, nil
}
