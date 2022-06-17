package telegram

import "go-telegram-bot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}

func New(c *telegram.Client) {

}
