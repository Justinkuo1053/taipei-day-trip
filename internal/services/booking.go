package services

// import (
// 	"fmt"
// 	"taipei-day-trip-go-go/internal/models"
// 	"taipei-day-trip-go-go/internal/repository/interfaces"
// )

// type BookingService struct {
// 	bookingRepo interfaces.BookingRepository
// 	orderRepo   interfaces.OrderRepository
// }

// func NewBookingService(
// 	bookingRepo interfaces.BookingRepository,
// 	orderRepo interfaces.OrderRepository,
// ) *BookingService {
// 	return &BookingService{
// 		bookingRepo: bookingRepo,
// 		orderRepo:   orderRepo,
// 	}
// }

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



