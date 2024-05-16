package controller

import (
	"github.com/shomali11/slacker"
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
