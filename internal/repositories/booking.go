package repositories

import (
	"taipei-day-trip-go-go/internal/models"

	"gorm.io/gorm"
)

type BookingRepository struct {
	DB *gorm.DB
}

func (r *BookingRepository) GetByUserID(userID uint) (*models.Booking, error) {
	var booking models.Booking
	err := r.DB.Preload("Attraction").Where("user_id = ?", userID).First(&booking).Error
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *BookingRepository) Create(booking *models.Booking) error {
	return r.DB.Create(booking).Error
}

func (r *BookingRepository) DeleteByUserID(userID uint) error {
	return r.DB.Where("user_id = ?", userID).Delete(&models.Booking{}).Error
}

func (r *BookingRepository) DeleteByID(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&models.Booking{}).Error
}
