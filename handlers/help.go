package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rumiani/gorate/lang"
	"github.com/rumiani/gorate/repository"
)

func MenuHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update, userRepo *repository.UserRepository) {
	userLang, err := userRepo.GetUserLanguage(update.Message.From.ID)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Something went wrong!"))
		return
	}
	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, lang.T(userLang, "menu")))
}
