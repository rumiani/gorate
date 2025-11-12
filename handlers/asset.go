package handlers

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rumiani/gorate/helpers"
	"github.com/rumiani/gorate/lang"
	"github.com/rumiani/gorate/repository"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func HandleAssetByCode(bot *tgbotapi.BotAPI, update tgbotapi.Update, repo *repository.AssetRepository, code string, userRepo *repository.UserRepository) {
	userLang, err := userRepo.GetUserLanguage(update.Message.From.ID)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Something went wrong!"))
		return
	}

	asset, err := repo.GetAssetByCode(code)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, lang.T(userLang, "notFound")))
		return
	}
	name := asset.EnName[0]
	if userLang == "fa" {
		name = asset.FaName[0]
	}

	// Format price with commas
	p := message.NewPrinter(language.English)
	commaPrice := p.Sprintf("%.0f", asset.CurrentPrice)

	msgText := fmt.Sprintf(
		"🔹%s\n💰 %s: %s\n🔗 /assets\n📜 /menu\n🆘 /help",
		helpers.Capitalize(name),
		lang.T(userLang, "price"),
		commaPrice,
	)

	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msgText))
}
