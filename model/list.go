package model

import (
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
)

func GetLists(api *anaconda.TwitterApi) ([]anaconda.List, error) {
	user, err := api.GetSelf(nil)
	id := user.Id
	if err != nil {
		return nil, err
	}
	lists, err := api.GetListsOwnedBy(id, nil)
	if err != nil {
		return nil, err
	}
	return lists, nil
}

func GetListStatuses(api *anaconda.TwitterApi, listID int64, count int) ([]StampTweet, error) {
	v := url.Values{}
	v.Set("count", strconv.Itoa(count))

	tweetList, err := api.GetListTweets(listID, true, v)
	if err != nil {
		return nil, err
	}

	stampTweetList, err := MakeStampTweet(tweetList)
	if err != nil {
		return nil, err
	}

	return stampTweetList, nil
}
