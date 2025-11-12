// handlers/assets.go
package handlers

import (
	"fmt"
	"strings"

	"github.com/rumiani/gorate/helpers"
	"github.com/rumiani/gorate/lang"

	"github.com/rumiani/gorate/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleAssets fetches all assets and sends a list of names and commands.
func HandleAssets(bot *tgbotapi.BotAPI, update tgbotapi.Update, repo *repository.AssetRepository, userRepo *repository.UserRepository) {
	userLang, err := userRepo.GetUserLanguage(update.Message.From.ID)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Something went wrong!"))
		return
	}

	assets, err := repo.GetAllAssets()
	if err != nil {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, lang.T(userLang, "failed")))
		return
	}

	types := map[string]string{}

	for _, asset := range assets {
		if _, ok := types[asset.Type]; !ok {
			types[asset.Type] = lang.T(userLang, asset.Type)
			types[asset.Type] += ":\n"
		}

		name := helpers.Capitalize(asset.EnName[0])
		if userLang == "fa" {
			name = asset.FaName[0]
		}
		command := strings.ToLower(asset.Code)
		types[asset.Type] += fmt.Sprintf("%s: /%s\n", name, command)
	}

	result := ""
	for _, t := range types {
		result += t + "\n"
	}
	result += "\n📜 /menu\n🆘 /help"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
	bot.Send(msg)
}
