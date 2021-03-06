package service

import (
	"github.com/meligGolang/domain"
	"fmt"
	"github.com/meligGolang/tweet_user"
	"os/user"
	"bufio"
	"os"
)

type TweetManager struct {
	users []*user.User
	tweets       []*domain.Tweet
	tweetsByUser map[string][]*domain.Tweet
}

func NewTweetManager() *TweetManager {

	tweetManager := new(TweetManager)

	tweetManager.users = make([]*user.User, 0)
	tweetManager.tweets = make([]*domain.Tweet, 0)
	tweetManager.tweetsByUser = make(map[string][]*domain.Tweet)

	return tweetManager
}
func (manager *TweetManager) registeredUser(user *tweet_user.User) bool {
	var found bool
	for !found && i < len(manager.users) {
		if len(manager.users) > 0 {
			if manager.users[i].Nick == user.Nick || manager.users[i].Mail == user.Mail {
				found = true
			} else {
				i++
			}
		} else {
			break
		}
	}
	return found
}
func (manager *TweetManager) PublishTweet(tweetToPublish *domain.Tweet) (int, error) {

	if tweetToPublish.User == nil {
		return 0, fmt.Errorf("user is required")
	}

	if tweetToPublish.Text == "" {
		return 0, fmt.Errorf("text is required")
	}

	if len(tweetToPublish.Text) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}

	if registredUser(tweetToPublish.User) {
		manager.users = append(manager.users, tweetToPublish.User)
	}
	manager.tweets = append(manager.tweets, tweetToPublish)

	tweetToPublish.Id = len(manager.tweets)

	//userTweets := manager.tweetsByUser[tweetToPublish.User]
	//manager.tweetsByUser[tweetToPublish.User] = append(userTweets, tweetToPublish)

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



func (manager *TweetManager) RegisterUser(user *tweet_user.User) {
	if !manager.registeredUser(user) {
		manager.users = append(manager.users, user)
	} else {
		fmt.Println("Usuario ya registrado")
	}
}

func (manager *TweetManager) GetUsers() []*tweet_user.User {
	var users []*tweet_user.User
	if manager.users != nil {
		users = manager.users
	}

	return users
}

func (manager *TweetManager) findTweet(id int) *domain.Tweet {
	i := 0
	var found bool
	var tweet *domain.Tweet

	for !found {
		if manager.tweets[i].Id == id {
			tweet = manager.tweets[i]
			return tweet
		} else {
			i++
		}
	}
	return nil
}

func (manager *TweetManager) EditTweet(id int) {
	input := bufio.NewReader(os.Stdin)
	tweet := manager.findTweet(id)
	var newText string
	fmt.Println("Editar tweet: ")
	newText, _ = input.ReadString('\n')
	tweet.Text = newText
}
