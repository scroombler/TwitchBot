package main

import (
	"fmt"

	"github.com/scroombler/TwitchBot/twitchbot"
)

func main() {
	channelName := "scroombler"
	fmt.Printf("Connecting to %v\n", channelName)

	creds, err := twitchbot.LoadCreds("./creds.json")
	if err != nil {
		fmt.Println("Error loading credentials")
		return
	}

	bot := twitchbot.New(creds)
	bot.Run()
}
