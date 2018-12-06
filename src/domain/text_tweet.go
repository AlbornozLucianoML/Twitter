package domain

import "time"

type TextTweet struct {
	User string
	Text string
	Date *time.Time
	Id int
}

func (textTweet *TextTweet) PrintableTweet() string {
	return "@" + textTweet.User + ": " + textTweet.Text
}

func (textTweet *TextTweet) GetId() int {
	return textTweet.Id
}

func (textTweet *TextTweet) SetId(id int) {
	textTweet.Id = id
}

func (textTweet *TextTweet) GetText() string {
	return textTweet.Text
}

func (textTweet *TextTweet) GetUser() string {
	return textTweet.User
}

func (TextTweet *TextTweet) GetDate() *time.Time {
	return TextTweet.Date
}

func NewTextTweet(user string, text string) *TextTweet {

	textTweet := TextTweet{User: user, Text: text}

	timeNow := time.Now()

	textTweet.User = user
	textTweet.Text = text
	textTweet.Date = &timeNow

	return &textTweet

}
