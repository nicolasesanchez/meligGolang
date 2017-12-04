package domain

import (
	"time"
	"meliGo/tweet_user"
)

type Tweet struct {
	User *tweet_user.User
	Text string
	Date *time.Time
	Id   int
}

func NewTweet(user *tweet_user.User, text string) *Tweet {

	date := time.Now()

	tweet := Tweet{
		User: user,
		Text: text,
		Date: &date,
	}

	return &tweet
}
