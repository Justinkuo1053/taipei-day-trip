package interfaces

import (
	"taipei-day-trip/internal/models"
	"time"
)

type BookingRepository interface {
	Create(booking *models.Booking) error                     // 創建預訂
	GetByID(id uint) (*models.Booking, error)                 // 查詢預訂詳情
	GetByUserID(userID uint) ([]models.Booking, error)        // 查詢用戶所有預訂
	Update(booking *models.Booking) error                     // 更新預訂資訊
	Delete(id uint) error                                     // 取消預訂
	CheckAvailability(attractionID uint, date time.Time) bool // 檢查是否可預訂
}
