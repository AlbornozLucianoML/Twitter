package domain

import "time"

type QuoteTweet struct {
	TextTweet
	QuotedTweet Tweet
}

func (quoteTweet *QuoteTweet) PrintableTweet() string {
	return "@" + quoteTweet.User + ": " + quoteTweet.Text + " \"@" + quoteTweet.QuotedTweet.GetUser() + ": " + quoteTweet.QuotedTweet.GetText() + "\""
}

func NewQuoteTweet(user string, text string, quotedTweet Tweet) *QuoteTweet {

	quoteTweet := QuoteTweet{TextTweet{User: user, Text: text}, quotedTweet}

	timeNow := time.Now()

	quoteTweet.User = user
	quoteTweet.Text = text
	quoteTweet.Date = &timeNow
	quoteTweet.QuotedTweet = quotedTweet

	return &quoteTweet

}