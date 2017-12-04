package domain

import (
	"time"
	"github.com/meligGolang/user"
	"fmt"
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

func (tweet *Tweet) PritableTweet() string{
	return fmt.Sprintf("%s  :  %s"; tweet.User , tweet.Text )
}

func (tweet * Tweet) String() string {

}


