package service

import "github.com/AlbornozLucianoML/Twitter/src/domain"

type TweetWriter interface {
	WriteTweet(domain.Tweet)
}