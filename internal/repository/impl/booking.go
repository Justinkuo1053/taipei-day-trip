package impl

import (
	"errors"
	"taipei-day-trip/internal/repository/interfaces"
	"time"
	"taipei-day-trip/internal/models"

	"gorm.io/gorm"
)

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) interfaces.BookingRepository {
	return &bookingRepository{db: db}
}

// 創建預訂
func (r *bookingRepository) Create(booking *models.Booking) error {
	// 檢查該時段是否已被預訂
	if !r.CheckAvailability(booking.AttractionID, booking.Date) {
		return errors.New("該時段已被預訂")
	}
	return r.db.Create(booking).Error
}

// 檢查可用性
func (r *bookingRepository) CheckAvailability(attractionID uint, date time.Time) bool {
	var count int64
	r.db.Model(&models.Booking{}).
		Where("attraction_id = ? AND date = ?", attractionID, date).
		Count(&count)
	return count == 0
}

// 獲取用戶所有預訂
func (r *bookingRepository) GetByUserID(userID uint) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Where("user_id = ?", userID).
		Preload("Attraction").
		Find(&bookings).Error
	return bookings, err
}
