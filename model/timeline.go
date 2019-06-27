package model

import (
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
)

func GetHomeTimeline(api *anaconda.TwitterApi, count int) ([]StampTweet, error) {
	v := url.Values{}
	v.Set("count", strconv.Itoa(count))

	tweetList, err := api.GetHomeTimeline(v)
	if err != nil {
		return nil, err
	}

	stampTweetList, err := MakeStampTweet(tweetList)
	if err != nil {
		return nil, err
	}

	return stampTweetList, nil

}

// 変換の都合上、userIDをstringのままで
func GetUserTimeline(api *anaconda.TwitterApi, userID string, count int) ([]StampTweet, error) {
	v := url.Values{}
	v.Set("count", strconv.Itoa(count))
	v.Set("user_id", userID)

	tweetList, err := api.GetUserTimeline(v)
	if err != nil {
		return nil, err
	}

	stampTweetList, err := MakeStampTweet(tweetList)
	if err != nil {
		return nil, err
	}

	return stampTweetList, nil
}
