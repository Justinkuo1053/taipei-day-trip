package interfaces

import "taipei-day-trip/internal/models"

type OrderRepository interface {
	Create(order *models.Order) error                      // 創建訂單
	GetByID(id uint) (*models.Order, error)                // 查詢訂單詳情
	GetByOrderNumber(number string) (*models.Order, error) // 依訂單號查詢
	GetByUserID(userID uint) ([]models.Order, error)       // 查詢用戶所有訂單
	Update(order *models.Order) error                      // 更新訂單狀態
	UpdatePaymentStatus(id uint, status string) error      // 更新支付狀態
}

type OrderService interface {
	GetOrderByID(id int) (*models.Order, error)
}
