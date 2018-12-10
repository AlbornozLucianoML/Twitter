package main

import (
	"github.com/AlbornozLucianoML/Twitter/src/application"
	"github.com/AlbornozLucianoML/Twitter/src/service"
)

func main() {

	tweetWriter := service.NewFileTweetWriter()

	tweetManager := service.NewTweetManager(tweetWriter)

	application.RunServer(tweetManager)

	application.RunShell(tweetManager)

}
