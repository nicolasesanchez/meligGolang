package main

import (
	"meliGo/service"
	"meliGo/tweet_user"
	"meliGo/domain"
	"fmt"
)

func main() {
	manager := service.NewTweetManager()

	var user *tweet_user.User
	var tweet *domain.Tweet

	user = tweet_user.NewUser("nico", "nico@asd.com", "mabel", "asd123")
	manager.RegisterUser(user)

	tweet = domain.NewTweet(user, "Hola mundo")

	i, _ := manager.PublishTweet(tweet)

	tweet = domain.NewTweet(user, "Hola mundo x2")

	i,_ = manager.PublishTweet(tweet)

	fmt.Println("Tweet 0:",manager.GetTweets()[0].Text)

	manager.EditTweet(1)

	fmt.Println(i)
	fmt.Println("users",len(manager.GetUsers()))
	fmt.Println("Tweet 0:",manager.GetTweets()[0].Text)


}
