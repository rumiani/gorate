package bot

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rumiani/gorate/lang"
	"github.com/rumiani/gorate/repository"
)

func CallbackQueryHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update, userService *repository.UserService) {
	data := update.CallbackQuery.Data
	chatID := update.CallbackQuery.Message.Chat.ID
	userID := update.CallbackQuery.From.ID

	if strings.HasPrefix(data, "set_lang") {
		langCode := data[9:] // "en" or "fa"
		// Only userID needed
		err := userService.SetUserLanguage(userID, langCode)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(chatID, "Something went wrong!"))
			fmt.Println(err)
			return
		}

		bot.Send(tgbotapi.NewMessage(chatID, lang.T(langCode, "welcome")+"\n/menu"))

		// Acknowledge callback
		bot.Request(tgbotapi.NewCallback(update.CallbackQuery.ID, lang.T(langCode, "langUpdated")))
	} else {
		bot.Send(tgbotapi.NewMessage(chatID, "Unknown callback query."))
	}
}
