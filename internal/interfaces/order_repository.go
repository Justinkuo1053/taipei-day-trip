package interfaces

import "taipei-day-trip-go-go/internal/models"

// 訂單資料存取介面，需實作所有方法
type OrderRepository interface {
	// 建立訂單
	CreateOrder(order *models.Order) error
	// 依 userID 與 bookingID 查詢訂單
	GetOrderByUserIDAndBookingID(userID, bookingID uint, order *models.Order) error
	// 依訂單編號查詢訂單（含景點與圖片）
	GetOrderByNumber(orderNumber string) (*models.Order, error)
}
