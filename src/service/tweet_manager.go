package service

import (
	"fmt"
	"github.com/AlbornozLucianoML/Twitter/src/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type TweetManager struct {
	tweet       domain.Tweet
	tweetsMap   map[string] []domain.Tweet
	Tweets      []domain.Tweet
	TweetWriter TweetWriter
}

func (tweetManager *TweetManager) PublishTweet(twit domain.Tweet) (int, error) {

	if twit.GetUser() == "" {
		return -1, fmt.Errorf("user is required")
	}

	if twit.GetText() == "" {
		return -1, fmt.Errorf("text is required")
	}

	if len(twit.GetText()) > 140 {
		return -1, fmt.Errorf("text length less than 140 characters is required")
	}

	id := len(tweetManager.Tweets)

	twit.SetId(id)

	tweetManager.tweet = twit

	if _, exists := tweetManager.tweetsMap[twit.GetUser()]; !exists {
		tweetManager.tweetsMap[twit.GetUser()] =  make([]domain.Tweet, 0)
	}

	tweetManager.tweetsMap[twit.GetUser()] = append(tweetManager.tweetsMap[twit.GetUser()], twit)

	tweetManager.Tweets = append(tweetManager.Tweets, twit)

	tweetManager.TweetWriter.WriteTweet(twit)

	return id, nil

}

func (tweetManager *TweetManager) GetTweet() domain.Tweet {

	return tweetManager.tweet

}

func (tweetManager *TweetManager) GetTweets() []domain.Tweet {

	return tweetManager.Tweets

}

func (tweetManager *TweetManager) GetTweetById(id int) domain.Tweet {

	return tweetManager.Tweets[id]

}

func (tweetManager *TweetManager) GetLastTweet() domain.Tweet {

	len := len(tweetManager.GetTweets())

	return tweetManager.Tweets[len - 1]

}

func (tweetManager *TweetManager) CountTweetsByUser(user string) int {

	count := 0

	for _, valor := range tweetManager.Tweets {
		if valor.GetUser() == user {
			count++
		}
	}

	return count
}

func (tweetManager *TweetManager) GetTweetsByUser(user string) []domain.Tweet {
	return tweetManager.tweetsMap[user]
}

func NewTweetManager(tweetWriter TweetWriter) *TweetManager {

	tweetsMap := make(map[string] []domain.Tweet)
	tweets := make([]domain.Tweet, 0)

	tweetManager := TweetManager{tweetsMap: tweetsMap, Tweets: tweets, TweetWriter: tweetWriter}

	return &tweetManager

}

func (tweetManager *TweetManager) SearchTweetsContaining(query string, searchResult chan domain.Tweet) {

	go func() { for _, valor := range tweetManager.Tweets {
		if strings.Contains(valor.GetText(), query) {
			searchResult <- valor
		}
	}
		close(searchResult)
	}()

}

func (tweetManager *TweetManager) PublishRest(c *gin.Context) {

	var tweetToPublish domain.TweetDTO

	if err := c.ShouldBindJSON(&tweetToPublish); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tweetPublished := GetTweet(&tweetToPublish, tweetManager)

	tweetManager.PublishTweet(tweetPublished)

	c.JSON(200, tweetPublished)

}

func (tweetManager *TweetManager) GetLastTweetRest(c *gin.Context) {

	c.JSON(200, tweetManager.GetLastTweet())

}

func (tweetManager *TweetManager) GetTweetsRest(c *gin.Context) {

	c.JSON(200, tweetManager.Tweets)

}

func (tweetManager *TweetManager) GetTweetByIdRest(c *gin.Context) {

	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	c.JSON(200, tweetManager.Tweets[idInt])

}


func (tweetManager *TweetManager) CountTweetsByUserRest(c *gin.Context) {

	user := c.Param("user")

	count := 0

	for _, valor := range tweetManager.Tweets {
		if valor.GetUser() == user {
			count++
		}
	}

	c.String(200, strconv.Itoa(count))

}

func (tweetManager *TweetManager) GetTweetsByUserRest(c *gin.Context) {

	user := c.Param("user")

	tweetsToPublish := tweetManager.tweetsMap[user]

	c.JSON(200, tweetsToPublish)

}

func (tweetManager *TweetManager) SearchTweetByQueryRest(c *gin.Context) {

	query := c.Param("query")

	searchResult := make(chan domain.Tweet)

	tweetManager.SearchTweetsContaining(query, searchResult)

	foundTweet := <- searchResult

	c.JSON(200, foundTweet)

}