package impl

import (
	"taipei-day-trip/internal/models" // 修正 import 路徑

	"gorm.io/gorm"
)

type attractionRepository struct {
	db *gorm.DB
}

// 創建景點
func (r *attractionRepository) Create(attraction *models.Attraction) error {
	return r.db.Create(attraction).Error
}

// 分頁查詢景點
func (r *attractionRepository) GetAll(page, limit int) ([]models.Attraction, error) {
	var attractions []models.Attraction
	offset := (page - 1) * limit

	err := r.db.Offset(offset).
		Limit(limit).
		Order("id desc").
		Find(&attractions).Error

	return attractions, err
}

// 搜尋景點
func (r *attractionRepository) Search(keyword string) ([]models.Attraction, error) {
	var attractions []models.Attraction

	err := r.db.Where("name LIKE ? OR description LIKE ?",
		"%"+keyword+"%", "%"+keyword+"%").
		Find(&attractions).Error

	return attractions, err
}
