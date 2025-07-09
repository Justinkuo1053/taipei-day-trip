package interfaces

import "taipei-day-trip-go-go/internal/models"

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	GetOrderByUserIDAndBookingID(userID, bookingID uint, order *models.Order) error
}
