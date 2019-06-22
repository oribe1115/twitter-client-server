package model

// 新規ツイート用
type NewTweet struct {
	Text string `json:"text,omitempty"`
}

func PostNewTweet(newTweet NewTweet) {
	tweet, err := api.PostTweet(newTweet.Text, nil)
}