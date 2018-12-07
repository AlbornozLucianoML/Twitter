package domain

import "time"

type ImageTweet struct {
	TextTweet
	UrlImage string
}

func (imageTweet *ImageTweet) PrintableTweet() string {
	return "@" + imageTweet.User + ": " + imageTweet.Text + " " + imageTweet.UrlImage
}

func NewImageTweet(user string, text string, url string) *ImageTweet {

	imageTweet := ImageTweet{TextTweet{User: user, Text: text}, url}

	timeNow := time.Now()

	imageTweet.User = user
	imageTweet.Text = text
	imageTweet.Date = &timeNow
	imageTweet.UrlImage = url

	return &imageTweet

}