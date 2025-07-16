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
	// 先預載 Attraction 物件
	err := r.DB.Preload("Attraction").Where("user_id = ?", userID).First(&booking).Error
	if err != nil {
		return nil, err
	}

	// 取得 Attraction 的所有圖片，塞進 Images 欄位
	var images []models.Image
	r.DB.Where("attraction_id = ?", booking.AttractionID).Find(&images)
	imageURLs := make([]string, 0, len(images))
	for _, img := range images {
		imageURLs = append(imageURLs, img.URL)
	}
	booking.Attraction.Images = imageURLs // <-- 這樣 handler 就能取得正確圖片網址

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
