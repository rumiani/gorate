package bot

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/rumiani/gorate/handlers"
	"github.com/rumiani/gorate/repository"
)

func CommandHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update, assetRepo *repository.AssetRepository, userRepo *repository.UserRepository) {

	switch update.Message.Text {
	case "/start":
		handlers.HandleStart(bot, update)
	case "/assets":
		handlers.HandleAssets(bot, update, assetRepo, userRepo)
	case "/menu":
		handlers.MenuHandler(bot, update, userRepo)
	default:
		str := update.Message.Text
		if strings.HasPrefix(str, "/") {
			str = str[1:]
			handlers.HandleAssetByCode(bot, update, assetRepo, str, userRepo)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command.")
			bot.Send(msg)
		}
	}
}
