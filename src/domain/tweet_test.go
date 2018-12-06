package domain

import "testing"

func TestCanGetAPrintableTweet(t *testing.T) {

	//Initialization
	tweet := NewTweet("grupoesfera", "This is my tweet")

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
	tweet := NewTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}