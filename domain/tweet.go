package domain

import (
	"time"
	"github.com/meligGolang/tweet_user"
	"fmt"
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

func (tweet *Tweet) PritableTweet() string{
	return fmt.Sprintf("%s  :  %s"; tweet.User , tweet.Text )
}

func (tweet * Tweet) String() string {

}


