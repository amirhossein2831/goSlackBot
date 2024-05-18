package controller

import "github.com/shomali11/slacker"

func AgeCommand(bot *slacker.Slacker) {
	bot.Command("i was born in <year>", &slacker.CommandDefinition{
		Description: "age calculator",
		Examples:    []string{"i was born in 2020"},
		Handler:     HandleAgeCommand,
	})

}

func WelcomeCommand(bot *slacker.Slacker) {
	bot.Command("Hi,i'm <name>", &slacker.CommandDefinition{
		Description: "say welcome to user",
		Examples:    []string{"Hi,i'm <name>"},
		Handler:     HandleWelcomeCommand,
	})

}

func WhatsUpCommand(bot *slacker.Slacker) {
	bot.Command("so what's up bro", &slacker.CommandDefinition{
		Description: "talk about the situation",
		Examples:    []string{"so what's up bro"},
		Handler:     HandleWhatUpCommand,
	})

}

func EOLWOFCommand(bot *slacker.Slacker) {
	bot.Command("say something to member of EOLWOF channel", &slacker.CommandDefinition{
		Description: "talk about some things",
		Examples:    []string{"say something to member of EOLWOF channel"},
		Handler:     HandleEOLWOFCommand,
	})

}

func KnowPeopleCommand(bot *slacker.Slacker) {
	bot.Command("do you know <name> <family>", &slacker.CommandDefinition{
		Description: "talk about people",
		Examples:    []string{"do you know mobin zaki"},
		Handler:     HandleKnowPeopleCommand,
	})
}

func KnowPeopleIICommand(bot *slacker.Slacker) {
	bot.Command("what about <name> <family>", &slacker.CommandDefinition{
		Description: "talk about people",
		Examples:    []string{"what about mobin zaki"},
		Handler:     HandleKnowPeopleCommand,
	})
}

func CheckValidDomainCommand(bot *slacker.Slacker) {
	bot.Command("is <domain> a valid domain", &slacker.CommandDefinition{
		Description: "check the validation of a domain",
		Examples:    []string{"is gmail.com a valid domain"},
		Handler:     HandleCheckValidDomainCommand,
	})
}
