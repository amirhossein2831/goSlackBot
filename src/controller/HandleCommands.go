package controller

import (
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"strconv"
)

func HandleAgeCommand(botCtx slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
	year := r.Param("year")
	yob, err := strconv.Atoi(year)
	if err != nil {
		log.Fatal(err)
	}
	age := 2024 - yob

	res := fmt.Sprintf("so your age is %v", age)

	err = w.Reply(res)
	if err != nil {
		log.Fatal(err)
	}

}
func HandleWelcomeCommand(botCtx slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
	name := r.Param("name")

	res := fmt.Sprintf("hi %v,welcome to our bot,i am goSlack and i'm here to help,what can i do for you?", name)

	err := w.Reply(res)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleWhatUpCommand(botCtx slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
	err := w.Reply("so, az zaman safavie : going the moon \n zaman ghajar : going the moon \n alanam ke : going the moon")
	if err != nil {
		log.Fatal(err)
	}
}

func HandleEOLWOFCommand(botCtx slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
	err := w.Reply("this people are the best in the universe, just i want to tell them life is short so enjoy every moment of your life")
	if err != nil {
		log.Fatal(err)
	}
}

func HandleKnowPeopleCommand(botCtx slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
	name := r.Param("name")
	family := r.Param("family")
	var res string
	if name == "mobin" && family == "zaki" {
		res = fmt.Sprintf("yes, %v %v is student of ce at aut, he is originaly from khoiy \n,"+
			"he has a good heart and help peolpe with nothing back ", name, family)
	} else if name == "sina" && family == "sadeghi" {
		res = fmt.Sprintf("he is the coolest guy in aut, also known as %v config", name)
	} else if name == "kian" && family == "poorazar" {
		res = fmt.Sprintf("loooooooool")
	} else {
		res = fmt.Sprintf("still need some info to know this people")
	}

	err := w.Reply(res)
	if err != nil {
		log.Fatal(err)
	}
}
