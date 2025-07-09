package implementations


import "taipei-day-trip-go-go/internal/interfaces"

// import (
// 	"errors"
// 	"taipei-day-trip-go-go/internal/interfaces"
// 	"taipei-day-trip-go-go/internal/models"
// 	"taipei-day-trip-go-go/internal/utils"
// )
=======
import (
	"fmt"
	"taipei-day-trip-go-go/internal/interfaces"
	"taipei-day-trip-go-go/internal/models"
	"taipei-day-trip-go-go/internal/utils"
)
>>>>>>> Stashed changes

type OrderServiceImpl struct{}

func NewOrderServiceImpl() interfaces.OrderService {
	return &OrderServiceImpl{}
}

func (s *OrderServiceImpl) CreateOrder(order models.OrderInput) (string, error) {
	if order.Prime != "test_prime" {
		return "", fmt.Errorf("付款失敗（mock）")
	}
	orderNumber := fmt.Sprintf("20250708%06d", 123456) // 可改為動態產生
	orderModel := models.Order{
		OrderNumber:  orderNumber,
		UserID:       1, // TODO: 之後從 session 或 token 取得
		BookingID:    1, // TODO: 之後根據實際 booking 流程取得
		AttractionID: order.Order.Trip.Attraction.ID,
		Price:        order.Order.Price,
		TripDate:     order.Order.Trip.Date,
		TripTime:     order.Order.Trip.Time,
		ContactName:  order.Order.Contact.Name,
		ContactEmail: order.Order.Contact.Email,
		ContactPhone: order.Order.Contact.Phone,
		Status:       1, // mock 直接設已付款
	}
	if err := utils.Database.Create(&orderModel).Error; err != nil {
		return "", fmt.Errorf("資料庫寫入失敗: %v", err)
	}
	return orderNumber, nil
}

func (s *OrderServiceImpl) GetOrder(orderNumber string) (*models.Order, error) {
	var order models.Order
	err := utils.Database.Preload("Attraction").Where("order_number = ?", orderNumber).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *OrderServiceImpl) ProcessPayment(orderNumber string, paymentData models.PaymentInput) error {
	return nil
}
