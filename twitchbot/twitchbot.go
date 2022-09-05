package twitchbot

import (
	"fmt"
	"net"
)

type TwitchBot struct {
	Creds *Cred
	Conn  net.Conn
}

const (
	url = "irc.chat.twitch.tv:6697"
)

func New(creds *Cred) *TwitchBot {
	return &TwitchBot{
		Creds: creds,
		Conn:  nil,
	}
}

func (b *TwitchBot) Shutdown() {
	// TODO disconnect from IRC
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
	passMsg := fmt.Sprintf("PASS oath:%v", b.Creds.Token)
	nickMsg := fmt.Sprintf("NICK ScroombleBot")

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

	return nil
}

func (b *TwitchBot) Run() {
	defer b.Shutdown()

	err := b.Connect() // TODO check error
	if err != nil {
		return
	}

	err = b.Authenticate()
	if err != nil {
		return
	}

	buf := make([]byte, 8192)

	for {
		n, err := b.Conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			// TODO reconnect logic
		}

		fmt.Println(buf[:n])
	}
}
