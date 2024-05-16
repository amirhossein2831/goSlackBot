package controller

import (
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"strconv"
)

func PrintCommandEvent(events <-chan *slacker.CommandEvent) {
	for event := range events {
		println("Command event is fired")
		println(event.Command)
		println(event.Parameters)
		println(event.Event)
		println("***********************")
	}
}

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
