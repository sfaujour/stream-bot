package bot

import (
	"log"
	"medgebot/irc"
)

// Prints messages to the console
func (bot *Bot) RegisterReadLogger() {
	bot.RegisterHandler(
		NewHandler(func(msg irc.Message) {
			log.Printf("> %s", msg.String())
		}),
	)
}
