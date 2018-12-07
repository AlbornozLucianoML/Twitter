package service

import "github.com/AlbornozLucianoML/Twitter/src/domain"

type MemoryTweetWriter struct {
	LastTweet domain.Tweet
}

func (memoryTweetWriter *MemoryTweetWriter) WriteTweet(tweet domain.Tweet) {
	memoryTweetWriter.LastTweet = tweet
}

func NewMemoryTweetWriter() *MemoryTweetWriter {

	memoryTweetWriter := MemoryTweetWriter{nil}

	return &memoryTweetWriter

}

func (memoryTweetWriter *MemoryTweetWriter) GetLastSavedTweet() domain.Tweet {
	return memoryTweetWriter.LastTweet
}