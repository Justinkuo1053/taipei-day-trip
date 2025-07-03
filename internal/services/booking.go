package services

import (
	"taipei-day-trip-go-go/internal/interfaces"
	"taipei-day-trip-go-go/internal/models"
)

type BookingService struct {
	BookingRepo interfaces.BookingRepository
}

func NewBookingService(bookingRepo interfaces.BookingRepository) *BookingService {
	return &BookingService{BookingRepo: bookingRepo}
}

// 透過使用者ID取得預訂
func (s *BookingService) GetBookingByUserID(userID uint) (*models.Booking, error) {
	return s.BookingRepo.GetByUserID(userID)
}

// 創建預訂
func (s *BookingService) CreateBooking(booking *models.Booking) error {
	return s.BookingRepo.Create(booking)
}

// 透過使用者ID刪除預訂
func (s *BookingService) DeleteBookingByUserID(userID uint) error {
	return s.BookingRepo.DeleteByUserID(userID)
}

// 根據訂單ID刪除預訂
func (s *BookingService) DeleteBookingByID(id uint) error {
	return s.BookingRepo.DeleteByID(id)
}

// // 創建預訂和訂單
// func (s *BookingService) CreateBookingWithOrder(
// 	booking *models.Booking,
// 	orderInfo *models.Order,
// ) error {
// 	// 創建預訂
// 	if err := s.bookingRepo.Create(booking); err != nil {
// 		return fmt.Errorf("創建預訂失敗: %w", err)
// 	}

// 	// 創建訂單
// 	orderInfo.BookingID = booking.ID
// 	if err := s.orderRepo.Create(orderInfo); err != nil {
// 		return fmt.Errorf("創建訂單失敗: %w", err)
// 	}

// 	return nil
// }
