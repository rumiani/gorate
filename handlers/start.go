// handlers/start.go
package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleStart sends a greeting to the user.
func HandleStart(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Choose your language 🔽 زبان خود را انتخاب کنید")

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🇬🇧 English", "set_lang_en"),
			tgbotapi.NewInlineKeyboardButtonData("🇮🇷 فارسی", "set_lang_fa"),
		),
	)

	msg.ReplyMarkup = buttons
	bot.Send(msg)
}
