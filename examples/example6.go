package main

import (
	"context"
	"github.com/shomali11/slacker"
	"log"
	"time"
)

func main() {
	bot := slacker.NewClient("<YOUR SLACK BOT TOKEN>")

	definition := &slacker.CommandDefinition{
		Description: "Server time!",
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			response.Typing()

			time.Sleep(time.Second)

			response.Reply(time.Now().Format(time.RFC1123))
		},
	}

	bot.Command("time", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
