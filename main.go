package main

import (
	"fmt"

	"github.com/SubhamShaww/go-discord-bot/bot"
	"github.com/SubhamShaww/go-discord-bot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
