package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TelegramID       string     `gorm:"column:telegramId;uniqueIndex;not null" json:"telegramId"`
	Name             *string    `json:"name,omitempty"`
	Username         *string    `gorm:"index" json:"username,omitempty"`
	IsBot            bool       `gorm:"column:isBot;default:false" json:"isBot"`
	LanguageCode     *string    `gorm:"column:languageCode" json:"languageCode,omitempty"`
	Status           string     `gorm:"type:varchar(20);default:'ACTIVE'" json:"status"`
	CreatedAt        time.Time  `gorm:"column:createdAt;autoCreateTime" json:"createdAt"`
	UpdatedAt        time.Time  `gorm:"column:updatedAt;autoUpdateTime" json:"updatedAt"`
	LastNotification *time.Time `gorm:"column:lastNotification" json:"lastNotification,omitempty"`
	NotificationPref string     `gorm:"column:notificationPref;type:varchar(20);default:'HOURLY'" json:"notificationPref"`
}

func (User) TableName() string {
	return `"User"`
}
