package twitchbot

import (
	"fmt"
	"net"
	"strings"
)

type TwitchBot struct {
	Creds *Cred
	Conn  net.Conn
}

const (
	url = "irc.chat.twitch.tv:6667"
)

func New(creds *Cred) *TwitchBot {
	return &TwitchBot{
		Creds: creds,
		Conn:  nil,
	}
}

func (b *TwitchBot) Shutdown() {
	b.SendChat("Scroomblebot leaving...")
	if b.Conn != nil {
		b.Conn.Close()
	}
}

func (b *TwitchBot) Connect() error {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		fmt.Println(err)
		return err
	}

	b.Conn = conn

	return nil
}

func (b *TwitchBot) Authenticate() error {
	passMsg := fmt.Sprintf("PASS oauth:%v\n", b.Creds.Token)
	nickMsg := fmt.Sprintf("NICK %v\n", b.Creds.Nickname)
	joinMsg := fmt.Sprintf("JOIN #%v\n", b.Creds.Channel)

	_, err := b.Conn.Write([]byte(passMsg))
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = b.Conn.Write([]byte(nickMsg))
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = b.Conn.Write([]byte(joinMsg))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (b *TwitchBot) SendChat(msg string) {
	if b.Conn == nil {
		return
	}

	cmd := fmt.Sprintf("PRIVMSG #%v :%v\n", b.Creds.Channel, msg)
	b.Conn.Write([]byte(cmd))
}

func (b *TwitchBot) Run() {
	defer b.Shutdown()

	err := b.Connect()
	if err != nil {
		return
	}

	err = b.Authenticate()
	if err != nil {
		return
	}

	buf := make([]byte, 8192)

	fmt.Println("Beginning main loop...")

	for {
		n, err := b.Conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from connection")
			fmt.Println(err)
			return
			// TODO reconnect logic
		}

		if n == 0 {
			return
		}

		fmt.Println(string(buf[:n]))
		go b.Handle(string(buf[:n]))
	}
}

func (b *TwitchBot) Handle(ircMsg string) {
	msg := strings.TrimSpace(strings.Split(ircMsg, ":")[2])
	fmt.Printf("Handling (%v)\n", msg)

	if msg == "!bot" {
		b.SendChat("Hello!")
	}
}
