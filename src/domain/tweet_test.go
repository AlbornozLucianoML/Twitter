package domain

import "testing"

func TestTextTweetPrintsUserAndTest(t *testing.T) {

	//Initialization
	tweet := NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}
}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	tweet := NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	tweet := NewImageTweet("grupoesfera", "This is my image",
		"http://www.grupoesfera.com.ar/common/img/grupoesfera.png")
	// Operation
	text := tweet.PrintableTweet()
	// Validation
	expectedText := "@grupoesfera: This is my image http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {
	// Initialization
	quotedTweet := NewTextTweet("grupoesfera", "This is my tweet")
	tweet := NewQuoteTweet("nick", "Awesome", quotedTweet)
	// Validation
	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	// Operation
	text := tweet.PrintableTweet()

	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}
}