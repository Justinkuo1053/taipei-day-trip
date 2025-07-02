package repositories

import (
	"taipei-day-trip-go-go/internal/models"

	"gorm.io/gorm"
)

type AttractionRepository struct {
	// 通常會包含資料庫連線或其他依賴
	DB *gorm.DB
}

// NewAttractionRepository 用來實體化 AttractionRepository
func NewAttractionRepository(db *gorm.DB) *AttractionRepository {
	return &AttractionRepository{DB: db}
}

func (r *AttractionRepository) GetAllAttractions() ([]models.Attraction, error) {
	// TODO: implement
	var attractions []models.Attraction
	return attractions, nil
}

func (r *AttractionRepository) GetByID(id uint) (*models.Attraction, error) {
	// TODO: implement
	var attraction *models.Attraction
	return attraction, nil
}

func (r *AttractionRepository) CreateAttraction(attraction *models.Attraction) error {
	// TODO: implement
	return nil
}

func (r *AttractionRepository) GetAll(page, limit int) ([]models.Attraction, error) {
	var attractions []models.Attraction
	offset := (page - 1) * limit
	result := r.DB.Limit(limit).Offset(offset).Find(&attractions)
	if result.Error != nil {
		return nil, result.Error
	}
	return attractions, nil
}

func (r *AttractionRepository) Search(keyword string) ([]models.Attraction, error) {
	var attractions []models.Attraction
	// ...查詢資料庫...
	return attractions, nil
}
