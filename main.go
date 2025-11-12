// main.go
package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	botCommands "github.com/rumiani/gorate/bot"
	"github.com/rumiani/gorate/db"
	"github.com/rumiani/gorate/lang"
	"github.com/rumiani/gorate/repository"
)

func main() {
	if err := lang.LoadTranslations(); err != nil {
		log.Fatalf("Error loading translations: %v", err)
	}

	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️ .env file not found, using system env variables")
	}
	env := os.Getenv("ENV") // dev or prod
	var token string
	if env == "prod" {
		token = os.Getenv("PROD_TOKEN")
	} else {
		token = os.Getenv("DEV_TOKEN")
	}
	if token == "" {
		log.Fatal("BOT_TOKEN not set")
	}

	// Initialize Telegram Bot API with the token from environment
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Optionally enable debug mode
	bot.Debug = false

	// Initialize database connection
	gormDB, err := db.InitDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	assetRepo := repository.NewAssetRepository(gormDB)
	userRepo := repository.NewUserRepository(gormDB)
	seviceRepo := repository.NewUserService(userRepo)

	// Set up update configuration
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates := bot.GetUpdatesChan(u)

	// Listen for updates
	for update := range updates {
		// Handle callback queries (button presses)
		if update.CallbackQuery != nil {
			botCommands.CallbackQueryHandler(bot, update, seviceRepo)
			continue
		}
		// Handle normal messages
		if update.Message != nil && update.Message.Text != "" {
			botCommands.CommandHandler(bot, update, assetRepo, userRepo)
		}
	}
}

/*
1. Watch ✅
2. Commas ✅
3. Translate ✅
4. Translate all
5. Github action
*/
