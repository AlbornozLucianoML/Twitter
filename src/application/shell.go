package application

import (
	"github.com/AlbornozLucianoML/Twitter/src/domain"
	"github.com/AlbornozLucianoML/Twitter/src/service"
	"github.com/abiosoft/ishell"
	"strconv"
	"time"
)

func RunShell(tweetManager *service.TweetManager) {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			c.Print("Write your user: ")

			user := c.ReadLine()

			tweet := domain.NewTextTweet(user, text)

			if _, err := tweetManager.PublishTweet(tweet); err != nil {
				c.Println("Error al publicar el tweet")
			}

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := tweetManager.GetTweet()

			c.Println(tweet.PrintableTweet())

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showMultipleTweets",
		Help: "Shows multiple tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tweetManager.GetTweets()

			for _, valor := range tweets {
				c.Println(valor.PrintableTweet())
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetId",
		Help: "Shows tweet's Id",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the Id: ")

			id := c.ReadLine()

			idInt, _ := strconv.Atoi(id)

			tweet := tweetManager.GetTweetById(idInt)

			c.Println("Usuario: ", tweet.GetUser(), " Tweet: ", tweet.GetText(), " Fecha: ", tweet.GetDate().Format(time.RFC850), "Id: ", tweet.GetId())

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countTweetsByUser",
		Help: "Count tweets by user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the user: ")

			user := c.ReadLine()

			tweetCounter := tweetManager.CountTweetsByUser(user)

			c.Println("User ", user, "tweeted ", tweetCounter, " tweets.")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetsByUser",
		Help: "Get tweets by user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the user: ")

			user := c.ReadLine()

			tweets := tweetManager.GetTweetsByUser(user)

			for _, valor := range tweets {
				c.Println("Usuario: ", valor.GetUser(), " Tweet: ", valor.GetText(), " Fecha: ", valor.GetDate().Format(time.RFC850), "Id: ", valor.GetId())
			}
			return
		},
	})

	shell.Run()

}
