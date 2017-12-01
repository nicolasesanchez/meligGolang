package service

import (
	"github.com/meligGolang/domain"
	"fmt"
	"go/ast"
)

type TweetManager struct {
	tweets       []*domain.Tweet
	tweetsByUser map[string][]*domain.Tweet
}

func NewTweetManager() *TweetManager {

	tweetManager := new(TweetManager)

	tweetManager.tweets = make([]*domain.Tweet, 0)
	tweetManager.tweetsByUser = make(map[string][]*domain.Tweet)

	return tweetManager
}

func (manager *TweetManager) PublishTweet(tweetToPublish *domain.Tweet) (int, error) {

	/*if tweetToPublish.User == "" {
		return 0, fmt.Errorf("user is required")
	}*/

	if tweetToPublish.Text == "" {
		return 0, fmt.Errorf("text is required")
	}

	if len(tweetToPublish.Text) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}

	manager.tweets = append(manager.tweets, tweetToPublish)

	tweetToPublish.Id = len(manager.tweets)

	userTweets := manager.tweetsByUser[tweetToPublish.User]
	manager.tweetsByUser[tweetToPublish.User] = append(userTweets, tweetToPublish)

	return tweetToPublish.Id, nil
}

// GetTweet returns the last published tweet
func (manager *TweetManager) GetTweet() *domain.Tweet {
	lastTweetIndex := len(manager.tweets) - 1

	return manager.tweets[lastTweetIndex]
}

func (manager *TweetManager) GetTweets() []*domain.Tweet {
	return manager.tweets
}

func (manager *TweetManager) GetTweetById(id int) *domain.Tweet {

	var tweet *domain.Tweet

	tweetIndex := 0

	for tweetIndex < len(manager.tweets) && tweet == nil {

		actualTweet := manager.tweets[tweetIndex]

		if actualTweet.Id == id {
			tweet = actualTweet
		}

		tweetIndex++
	}

	return tweet
}

func (manager *TweetManager) CountTweetsByUser(user string) int {

	var count int

	for _, tweet := range manager.tweets {
		if tweet.User == user {
			count++
		}
	}

	return count
}

func (manager *TweetManager) GetTweetsByUser(user string) []*domain.Tweet {

	return manager.tweetsByUser[user]
}

func  (manager *TweetManager) DeleteTweet(id int) string  {
	tweet := manager.GetTweetById(id)


}



