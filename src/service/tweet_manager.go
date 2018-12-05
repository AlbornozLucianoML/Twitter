package service

var Tweet = ""

func PublishTweet(tweet string) {

	Tweet = tweet

}

func GetTweet() string {
	return Tweet
}
