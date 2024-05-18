package main

import (
	"context"
	"github.com/amirhossein2831/goSlackBot/src/controller"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	"log"
	"os"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	SBToken := os.Getenv("SLACK_BOT_TOKEN")
	SAToken := os.Getenv("SLACK_APP_TOKEN")

	// instantiate the bot
	bot := slacker.NewClient(SBToken, SAToken)

	// define the Commands below
	controller.AgeCommand(bot)
	controller.WelcomeCommand(bot)
	controller.WhatsUpCommand(bot)
	controller.EOLWOFCommand(bot)
	controller.KnowPeopleCommand(bot)
	controller.KnowPeopleIICommand(bot)

	// log the running commands
	go controller.PrintCommandEvent(bot.CommandEvents())

	// defer and close the context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start serve the bot
	err = bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
