package application

import (
	"github.com/AlbornozLucianoML/Twitter/src/service"
	"github.com/gin-gonic/gin"
)

func RunServer(tweetManager *service.TweetManager) {

	router := gin.Default()

	router.GET("/getTweetList", tweetManager.GetTweetsRest)
	router.GET("/getTweet/:id", tweetManager.GetTweetsByIdRest)
	router.GET("/getLastTweet", tweetManager.GetLastTweetRest)


	go router.Run()

}