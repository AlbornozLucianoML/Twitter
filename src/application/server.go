package application

import (
	"github.com/AlbornozLucianoML/Twitter/src/service"
	"github.com/gin-gonic/gin"
)



func RunServer(tweetManager *service.TweetManager) {

	router := gin.Default()

	router.GET("/getTweetList", tweetManager.GetTweetsRest)
	router.GET("/getTweet/:id", tweetManager.GetTweetByIdRest)
	router.GET("/getLastTweet", tweetManager.GetLastTweetRest)
	router.POST("/publishTweet", tweetManager.PublishRest)
	router.GET("/countTweetsByUser/:user", tweetManager.CountTweetsByUserRest)
	router.GET("/getTweetsByUser/:user", tweetManager.GetTweetsByUserRest)
	router.GET("/searchTweetByQuery/:query", tweetManager.SearchTweetByQueryRest)

	go router.Run()

}