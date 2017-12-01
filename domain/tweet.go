package domain

import (
	"time"
	"meliGo/user"
)

type Tweet struct {
	User *user.User
	Text string
	Date *time.Time
	Id int
}

func NewTweet(user user.User, text string) *Tweet {

	date := time.Now()

	tweet := Tweet{
		User: &user,
		Text: text,
		Date: &date,
	}

	return &tweet
}



