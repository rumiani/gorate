package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Asset struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Code         string         `json:"code"`
	EnName       pq.StringArray `gorm:"column:enName;type:text[]" json:"enName"`
	FaName       pq.StringArray `gorm:"column:faName;type:text[]" json:"faName"`
	BuyCode      string         `json:"buyCode"`
	SellCode     string         `json:"sellCode"`
	Type         string         `json:"type"`
	CurrentPrice float64        `json:"currentPrice"`
	Status       string         `json:"status"`
	UpdatedAt    string         `json:"updatedAt"`
}

// Tell GORM real table name
func (Asset) TableName() string {
	return `"Asset"`
}
