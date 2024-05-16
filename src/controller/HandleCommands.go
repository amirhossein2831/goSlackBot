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

	res := fmt.Sprintf("hi %v,welcome to our bot,i am goSlack and i'm here to helo,what can i do for you?", name)

	err := w.Reply(res)
	if err != nil {
		log.Fatal(err)
	}
}
