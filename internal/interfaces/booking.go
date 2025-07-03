package interfaces

import "taipei-day-trip-go-go/internal/models"

type BookingRepository interface {
	GetByUserID(userID uint) (*models.Booking, error)
	Create(booking *models.Booking) error
	DeleteByUserID(userID uint) error
	DeleteByID(id uint) error
}
