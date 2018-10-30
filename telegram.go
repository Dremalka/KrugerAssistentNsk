package main

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)


func start(cfg Config) {
	bot, err := tgbotapi.NewBotAPI(cfg.TokenTelegram)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		
		if update.Message == nil {
			continue
		}

		if update.Message.Chat.Title == cfg.TittleChat {
			log.Printf("Leave chat: %s", update.Message.Chat)
			bot.LeaveChat(update.Message.Chat.ChatConfig())
		}

		//log.Printf("Chat: %s", update.Message.Chat)

		//log.Printf("[%s] %s", update.Message.From.ID, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
