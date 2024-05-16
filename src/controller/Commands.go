package controller

import "github.com/shomali11/slacker"

func AgeCommand(bot *slacker.Slacker) {
	bot.Command("i was born in <year>", &slacker.CommandDefinition{
		Description: "age calculator",
		Examples:    []string{"i was born in 2020"},
		Handler:     HandleAgeCommand,
	})

}
