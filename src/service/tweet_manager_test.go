package service_test

import (
	"github.com/AlbornozLucianoML/Twitter/src/domain"
	"github.com/AlbornozLucianoML/Twitter/src/service"
	"testing"
)

func IsValidTweet(t *testing.T, publishedTest domain.Tweet, id int, user string, text string) bool {

	if (publishedTest.GetUser() != user) {
		return false
	}

	if (publishedTest.GetText() != text) {
		return false
	}

	if (publishedTest.GetId() != id) {
		return false
	}

	return true

}

func TestPublishedTweetIsSaved(t *testing.T) {

	tweetManager := service.NewTweetManager()
	// Initialization
	var tweet domain.Tweet
	user := "mercadolibre"
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	//Operation
	if _, err := tweetManager.PublishTweet(tweet); err != nil {
		t.Error("Error not expected: ", err.Error())
	}

	//Validation

	publishedTweet := tweetManager.GetTweet()
	if publishedTweet.GetUser() != user || publishedTweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.GetUser(), publishedTweet.GetText())
	}

	if publishedTweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	//Initialization
	tweetManager := service.NewTweetManager()

	var tweet domain.Tweet

	var text string
	user := "User"

	tweet = domain.NewTextTweet(user, text)

	//Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	//Validation
	if err == nil || err.Error() != "text is required" {
		t.Error("Expected error is text is required, received: ", err.Error())
	}

}

func TestTweetWhichExceeding140CharacterIsNotPublised(t *testing.T) {

	//Initialization
	tweetManager := service.NewTweetManager()

	var tweet domain.Tweet

	user := "usuario"
	text1 := "0123456789" +
		"0123456789" +
		"0123456789" +
		"0123456789" +
		"0123456789" +
		"0123456789" +
		"0123456789" +
		"0123456789" +
		"0123456789" +
		"0123456789" +
		"0123456789" +
		"0123456789" +
		"0123456789" +
		"0123456789" +
		"0"

	tweet = domain.NewTextTweet(user, text1)

	//Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	//Validation
	if err == nil || err.Error() != "text length less than 140 characters is required" {
		t.Error("Expected error is text length less than 140 characters is required, received: ", err.Error())
	}

}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization

	tweetManager := service.NewTweetManager()

	var tweet, secondTweet domain.Tweet // Fill the tweets with data

	userPrimerTweet := "Luciano"
	textPrimerTweet := "Primer Tweet"

	userSegundoTweet := "Albornoz"
	textSegundoTweet := "Segundo Tweet"

	tweet = domain.NewTextTweet(userPrimerTweet, textPrimerTweet)

	secondTweet = domain.NewTextTweet(userSegundoTweet, textSegundoTweet)

	// Operation
	if _, err := tweetManager.PublishTweet(tweet); err != nil  {
		t.Error("Error not expected, got: ", err.Error())
	}

	if _, err := tweetManager.PublishTweet(secondTweet); err != nil  {
		t.Error("Error not expected, got: ", err.Error())
	}

	// Validation
	publishedTweets := tweetManager.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}
	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]
	if !IsValidTweet(t, firstPublishedTweet, 0, userPrimerTweet, textPrimerTweet) {
		return
	}

	if !IsValidTweet(t, secondPublishedTweet, 1, userSegundoTweet, textSegundoTweet) {
		return
	}
}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation

	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	IsValidTweet(t, publishedTweet, id, user, text)

}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	if _, err := tweetManager.PublishTweet(tweet); err != nil {
		return
	}
	if _, err := tweetManager.PublishTweet(secondTweet); err != nil {
		return
	}
	if _, err := tweetManager.PublishTweet(thirdTweet); err != nil {
		return
	}

	// Operation
	count := tweetManager.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)
	// publish the 3 tweets

	if _, err := tweetManager.PublishTweet(tweet); err != nil {
		return
	}
	if _, err := tweetManager.PublishTweet(secondTweet); err != nil {
		return
	}
	if _, err := tweetManager.PublishTweet(thirdTweet); err != nil {
		return
	}

	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 { /* handle error */}
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	// check if isValidTweet for firstPublshiedTweet and secondPublishedTweet
	if !IsValidTweet(t, firstPublishedTweet, 0, user, text) {
		return
	}

	if !IsValidTweet(t, secondPublishedTweet, 1, user, secondText) {
		return
	}
}


