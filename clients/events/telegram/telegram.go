package telegram

import (
	"go-telegram-bot/clients/telegram"
	"go-telegram-bot/lib/storage"
)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

func New(c *telegram.Client) {

}
