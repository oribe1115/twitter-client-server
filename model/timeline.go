package model

import (
	"net/url"
	"strconv"
)

func GetHomeTimeline(count int) ([]StampTweet, error) {
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
