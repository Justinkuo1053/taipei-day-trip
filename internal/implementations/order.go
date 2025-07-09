package implementations

import (
	"fmt"
	"math/rand"
	"taipei-day-trip-go-go/internal/interfaces"
	"taipei-day-trip-go-go/internal/models"
	"taipei-day-trip-go-go/internal/utils"
	"time"
)

type OrderServiceImpl struct{}

func NewOrderServiceImpl() interfaces.OrderService {
	return &OrderServiceImpl{}
}

func (s *OrderServiceImpl) CreateOrder(order models.OrderInput) (string, error) {
	if order.Prime != "test_prime" {
		return "", fmt.Errorf("付款失敗（mock）")
	}
	userID := uint(1)    // TODO: 之後從 session 或 token 取得
	bookingID := uint(1) // TODO: 之後根據實際 booking 流程取得
	// 防止同一 user、同一 booking 重複下訂
	var existOrder models.Order
	err := utils.Database.Where("user_id = ? AND booking_id = ?", userID, bookingID).First(&existOrder).Error
	if err == nil {
		return "", fmt.Errorf("此預約已建立訂單，請勿重複下訂")
	}
	rand.Seed(time.Now().UnixNano())
	orderNumber := fmt.Sprintf("%s%06d", time.Now().Format("20060102150405"), rand.Intn(1000000))
	orderModel := models.Order{
		OrderNumber:  orderNumber,
		UserID:       userID,
		BookingID:    bookingID,
		AttractionID: uint(order.Order.Trip.Attraction.ID),
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
