package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date *time.Time
	Id int
}

func (tweet Tweet) String() string {
	return tweet.PrintableTweet()
}

func (tweet *Tweet) PrintableTweet() string {
	return "@" + tweet.User + ": " + tweet.Text
}

func NewTweet(user string, text string) *Tweet {

	tweet := Tweet{User: user, Text: text}

	timeNow := time.Now()

	tweet.User = user
	tweet.Text = text
	tweet.Date = &timeNow

	return &tweet

}
