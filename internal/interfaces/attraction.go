package interfaces

import "taipei-day-trip-go-go/internal/models"

type AttractionRepository interface {
	GetByID(id uint) (*models.Attraction, error)
	GetAll(page, limit int) ([]models.Attraction, error)
	Search(keyword string) ([]models.Attraction, error)
}

type AttractionService interface {
	GetAttractionByID(id int) (*models.Attraction, error)
	ListAttractions(page, limit int) ([]models.Attraction, error)
	SearchAttractions(keyword string) ([]models.Attraction, error)
}
