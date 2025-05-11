package interfaces

import "taipei-day-trip/internal/models"

type AttractionRepository interface {
	Create(attraction *models.Attraction) error
	GetByID(id uint) (*models.Attraction, error)
	GetAll(page, limit int) ([]models.Attraction, error)
	GetByCategory(category string) ([]models.Attraction, error)
	Update(attraction *models.Attraction) error
	Delete(id uint) error
	Search(keyword string) ([]models.Attraction, error) // 新增 Search 方法
}

type AttractionService interface {
	GetAttractionByID(id int) (*models.Attraction, error)
}
