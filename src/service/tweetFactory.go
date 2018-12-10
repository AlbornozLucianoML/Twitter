package service

import (
	"github.com/AlbornozLucianoML/Twitter/src/domain"
)

func GetTweet(tweet *domain.TweetDTO, tweetManager *TweetManager) domain.Tweet {

	var newTweet domain.Tweet

	if tweet.Id == nil {
		if tweet.Url == "" {
			newTweet = domain.NewTextTweet(tweet.User, tweet.Text)
		} else {
			newTweet = domain.NewImageTweet(tweet.User, tweet.Text, tweet.Url)
		}
	} else {
		newTweet = domain.NewQuoteTweet(tweet.User, tweet.Text,
			tweetManager.Tweets[*tweet.Id])
	}

	return newTweet
}