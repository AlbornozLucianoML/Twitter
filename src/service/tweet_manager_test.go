package service_test

import (
	"github.com/AlbornozLucianoML/Twitter/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	tweet := "This is my first tweet"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}
}