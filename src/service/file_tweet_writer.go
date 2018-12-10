package service

import (
	"github.com/AlbornozLucianoML/Twitter/src/domain"
	"os"
)

type FileTweetWriter struct {
	file *os.File
}

func NewFileTweetWriter() *FileTweetWriter {

	file, _ := os.OpenFile(
		"Tweets.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)

	writer := new(FileTweetWriter)
	writer.file = file

	return writer
}

func (writer *FileTweetWriter) WriteTweet(tweet domain.Tweet) {

	go func() {
		if writer.file != nil {
			byteSlice := []byte(tweet.PrintableTweet() + "\n")
			writer.file.Write(byteSlice)
		}
	}()
}