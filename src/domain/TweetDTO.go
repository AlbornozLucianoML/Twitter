package domain

import "time"

type TweetDTO struct {
	User string
	Text string
	Date *time.Time
	Id int
}