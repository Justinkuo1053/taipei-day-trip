package services

import (
	"fmt"
	"math/rand"
	"taipei-day-trip/internal/models"     // 修正 import 路徑
	"taipei-day-trip/internal/repository" // 修正 import 路徑
	"taipei-day-trip/internal/repository/interfaces"
	"time"
)

type orderService struct {
	repo   repository.OrderRepository
	tappay TapPayClient
}

func NewOrderService(repo repository.OrderRepository, tappay TapPayClient) interfaces.OrderService {
	return &orderService{
		repo:   repo,
		tappay: tappay,
	}
}

func (s *orderService) CreateOrder(order *models.Order) error {
	// 生成訂單編號
	order.OrderNumber = generateOrderNumber()

	if err := s.repo.Create(order); err != nil {
		return fmt.Errorf("建立訂單失敗: %w", err)
	}
	return nil
}

func (s *orderService) ProcessPayment(orderID uint, prime string) error {
	order, err := s.repo.GetByID(orderID)
	if err != nil {
		return fmt.Errorf("取得訂單失敗: %w", err)
	}

	// 呼叫 TapPay API
	paymentResult, err := s.tappay.ProcessPayment(prime, order.TotalAmount, order.OrderNumber)
	if err != nil {
		return fmt.Errorf("處理付款失敗: %w", err)
	}

	// 更新訂單狀態
	if paymentResult.Status == "success" {
		if err := s.repo.UpdatePaymentStatus(orderID, "paid"); err != nil {
			return fmt.Errorf("更新付款狀態失敗: %w", err)
		}
	}

	return nil
}

func generateOrderNumber() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("ORD%s%d",
		time.Now().Format("20060102"),
		rand.Intn(10000))
}

package impl

type TapPayClient struct {
	// 假設需要 API 金鑰
	APIKey string
}

// 新增模擬方法
func (c *TapPayClient) ProcessPayment(orderID int, amount float64) error {
	// 模擬付款邏輯
	return nil
}
