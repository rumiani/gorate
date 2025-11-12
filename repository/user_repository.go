package repository

import (
	"strconv"

	"github.com/rumiani/gorate/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateOrUpdateLanguage stores the language for a user
func (r *UserRepository) CreateOrUpdateLanguage(userID int64, lang string) error {
	var user models.User
	tID := strconv.FormatInt(userID, 10) // convert Telegram ID to string

	// query by TelegramID, not by UUID primary key
	result := r.DB.First(&user, `"telegramId" = ?`, tID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			user = models.User{
				TelegramID:   tID,
				LanguageCode: &lang,
			}
			return r.DB.Create(&user).Error
		}
		return result.Error
	}

	user.LanguageCode = &lang
	return r.DB.Save(&user).Error
}

// GetUserLanguage gets the language from DB
func (r *UserRepository) GetUserLanguage(userID int64) (string, error) {
	var user models.User
	tID := strconv.FormatInt(userID, 10) // convert Telegram ID to string

	// query by TelegramID, not by UUID primary key
	if err := r.DB.First(&user, `"telegramId" = ?`, tID).Error; err != nil {
		return "en", err
	}

	if user.LanguageCode != nil {
		return *user.LanguageCode, nil
	}
	return "en", nil
}
