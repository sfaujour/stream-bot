package bot

import (
	"fmt"
	"log"
	"medgebot/irc"
)

type Bot struct {
	client *irc.Irc
	// TODO maybe a channel of chat lines?
}

func New() Bot {
	client := irc.NewClient()
	return Bot{
		client: client,
	}
}

// Connect to the bot client
func (bot *Bot) Connect() error {
	if err := bot.client.Connect("wss", "irc-ws.chat.twitch.tv:443"); err != nil {
		log.Printf("ERROR: bot connect - %s", err)
		return err
	}

	// make(chan irc.Message)
	// go readChat(messageChan)
	// go pingPong(messageChan)

	return nil
}

// Close the connection to the client
func (bot *Bot) Close() {
	bot.client.Close()
}

// Authenticate connects to the IRC stream with the given nick and password
func (bot *Bot) Authenticate(nick, password string) error {
	if err := bot.client.SendPass(password); err != nil {
		log.Printf("ERROR: send PASS failed: %s", err)
		return err
	}
	log.Println("< PASS ***")

	if err := bot.client.SendNick(nick); err != nil {
		log.Printf("ERROR: send NICK failed: %s", err)
		return err
	}

	return nil
}

// Join joins to a specific channel on the IRC
func (bot *Bot) Join(channel string) error {
	err := bot.client.Join(channel)
	return err
}

func (bot *Bot) ReadStreamFunc(work func(string)) {
	go func() {
		for {
			msg, err := bot.client.Read()
			if err != nil {
				log.Println("ERROR: read - " + err.Error())
				break
			}
			work(msg.String())
		}
	}()
}

// func (bot *Bot) ReadStream(stream chan<- string) {
// go func(stream chan<- string) {

// }(stream)
// }

// PrivMsg sends a message to the given channel, without prefix
func (bot *Bot) SendMessage(channel, message string) string {
	return fmt.Sprintf("PRIVMSG #%s %s", channel, message)
}

// readChat Reads from the client and passes the parsed messages to the stream channel
func (bot *Bot) readChat(stream chan<- irc.Message) {
	for {
		msg, err := bot.client.Read()
		if err != nil {
			log.Println("ERROR: read - " + err.Error())
			break
		}
		stream <- msg
	}
}

// PingPong responds to Ping heartbeats from the server so we don't get disconnected
func (bot *Bot) pingPong(stream chan irc.Message) {
	select {
	case msg := <-stream:
		if msg.Command == "PING" {
			pong := irc.Message{
				Command: "PONG",
				Params:  msg.Params,
			}

			if err := bot.client.Write(pong); err != nil {
				log.Printf("ERROR: send PONG failed: %s", err)
			}
		}
	}
}
