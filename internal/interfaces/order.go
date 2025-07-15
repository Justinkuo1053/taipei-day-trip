package interfaces

import "taipei-day-trip-go-go/internal/models"

type OrderService interface {
	CreateOrder(order models.OrderInput, userID uint) (string, error)
	GetOrder(orderNumber string) (*models.Order, error)
	ProcessPayment(orderNumber string, paymentData models.PaymentInput) error
}
