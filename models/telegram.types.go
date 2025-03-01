package models

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// CommandsDto -> dto for bot base command as </start>, etc.
type CommandsDto struct {
	Bot      *tgbotapi.BotAPI
	UserName string
	ChatId   int64
}
