package domain

import "time"

type Tweet interface {
	PrintableTweet() string
	GetId() int
	GetUser() string
	GetDate() *time.Time
	GetText() string
	SetId(int)
}