// repository/asset_repository.go
package repository

import (
	"github.com/rumiani/gorate/models"

	"gorm.io/gorm"
)

type AssetRepository struct {
	DB *gorm.DB
}

// NewAssetRepository creates a new AssetRepository with the given GORM DB.
func NewAssetRepository(db *gorm.DB) *AssetRepository {
	return &AssetRepository{DB: db}
}

// GetAllAssets retrieves all Asset records from the database.
func (r *AssetRepository) GetAllAssets() ([]models.Asset, error) {
	var assets []models.Asset
	// result := r.DB.Find(&assets)
	result := r.DB.Select("code", "enName", "faName", "type").Find(&assets)

	return assets, result.Error
}

// GetAssetByCode retrieves a single Asset from the database by its code.
func (r *AssetRepository) GetAssetByCode(code string) (*models.Asset, error) {
	var asset models.Asset
	result := r.DB.Where("code = ?", code).First(&asset)
	if result.Error != nil {
		return nil, result.Error
	}
	return &asset, nil
}
