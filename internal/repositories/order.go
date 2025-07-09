package repositories

import (
	"taipei-day-trip-go-go/internal/models"
	"taipei-day-trip-go-go/internal/utils"
)

type OrderRepositoryImpl struct{}

func NewOrderRepositoryImpl() *OrderRepositoryImpl {
	return &OrderRepositoryImpl{}
}

func (r *OrderRepositoryImpl) CreateOrder(order *models.Order) error {
	return utils.Database.Create(order).Error
}

func (r *OrderRepositoryImpl) GetOrderByUserIDAndBookingID(userID, bookingID uint, order *models.Order) error {
	return utils.Database.Where("user_id = ? AND booking_id = ?", userID, bookingID).First(order).Error
}
