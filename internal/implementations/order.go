package implementations

import (
	"fmt"
	"taipei-day-trip-go-go/internal/interfaces"
	"taipei-day-trip-go-go/internal/models"
)

type OrderServiceImpl struct{}

func NewOrderServiceImpl() interfaces.OrderService {
	return &OrderServiceImpl{}
}

func (s *OrderServiceImpl) CreateOrder(order models.OrderInput) (string, error) {
	if order.Prime == "test_prime" {
		orderNumber := fmt.Sprintf("20250708%06d", 123456) // 假編號
		return orderNumber, nil
	}
	return "", fmt.Errorf("付款失敗（mock）")
}

func (s *OrderServiceImpl) GetOrder(orderNumber string) (*models.Order, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *OrderServiceImpl) ProcessPayment(orderNumber string, paymentData models.PaymentInput) error {
	return nil
}
