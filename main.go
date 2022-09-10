package main

import (
	"fmt"

	"github.com/scroombler/TwitchBot/twitchbot"
)

func main() {
	creds, err := twitchbot.LoadCreds("./creds.json")
	if err != nil {
		fmt.Println("Error loading credentials")
		return
	}

	bot := twitchbot.New(creds)
	bot.Run()
}

/*
Chatbot flow

1. Connect to the Twitch endpoint (e.g. irc.chat.twitch.tv)
  a. Open a TCP socket, maintain a connection
2. Authenticate ourself (send PASS, NICK)
3. Join a server
3. Main loop
  a. Receive message from the chat room
  b. Respond to commands
  c. Heartbeat
*/
