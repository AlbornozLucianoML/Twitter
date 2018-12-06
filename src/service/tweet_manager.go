package service

import (
	"fmt"
	"github.com/AlbornozLucianoML/Twitter/src/domain"
)

type TweetManager struct {
	tweet domain.Tweet
	tweetsMap map[string] []domain.Tweet
	tweets []domain.Tweet
}

func (tweetManager *TweetManager) PublishTweet(twit domain.Tweet) (int, error) {

	if twit.GetUser() == "" {
		return -1, fmt.Errorf("user is required")
	}

	if twit.GetText() == "" {
		return -1, fmt.Errorf("text is required")
	}

	if len(twit.GetText()) > 140 {
		return -1, fmt.Errorf("text length less than 140 characters is required")
	}

	id := len(tweetManager.tweets)

	twit.SetId(id)

	tweetManager.tweet = twit

	if _, existe := tweetManager.tweetsMap[twit.GetUser()]; !existe {
		tweetManager.tweetsMap[twit.GetUser()] =  make([]domain.Tweet, 0)
	}

	tweetManager.tweetsMap[twit.GetUser()] = append(tweetManager.tweetsMap[twit.GetUser()], twit)

	tweetManager.tweets = append(tweetManager.tweets, twit)

	return id, nil

}

func (tweetManager *TweetManager) GetTweet() domain.Tweet {

	return tweetManager.tweet

}

func (tweetManager *TweetManager) GetTweets() []domain.Tweet {

	return tweetManager.tweets

}

func (tweetManager *TweetManager) GetTweetById(id int) domain.Tweet {

	return tweetManager.tweets[id]

}

//func (tweetManager TweetManager) InitializeService() {
//
//	tweetManager.tweetsMap = make(map[string] []*domain.Tweet)
//	tweetManager.tweets = make([]*domain.Tweet, 0)
//
//}

func (tweetManager *TweetManager) CountTweetsByUser(user string) int {

	count := 0

	for _, valor := range tweetManager.tweets {
		if valor.GetUser() == user {
			count++
		}
	}

	return count
}

func (tweetManager *TweetManager) GetTweetsByUser(user string) []domain.Tweet {
	return tweetManager.tweetsMap[user]
}

func NewTweetManager() *TweetManager {

	tweetsMap := make(map[string] []domain.Tweet)
	tweets := make([]domain.Tweet, 0)

	tweetManager := TweetManager{tweetsMap: tweetsMap, tweets: tweets}

	return &tweetManager

}
