package impl

import (
	"taipei-day-trip/internal/models"
	"taipei-day-trip/internal/repository/interfaces"

	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &orderRepository{db: db}
}

// 創建訂單
func (r *orderRepository) Create(order *models.Order) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 創建訂單
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		// 更新相關預訂狀態
		if err := tx.Model(&models.Booking{}).
			Where("id = ?", order.BookingID).
			Update("status", "booked").Error; err != nil {
			return err
		}

		return nil
	})
}

// 更新支付狀態
func (r *orderRepository) UpdatePaymentStatus(id uint, status string) error {
	return r.db.Model(&models.Order{}).
		Where("id = ?", id).
		Update("payment_status", status).Error
}

// 獲取用戶所有訂單
func (r *orderRepository) GetByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("user_id = ?", userID).
		Preload("Booking").
		Preload("Booking.Attraction").
		Find(&orders).Error
	return orders, err
}

// 新增 GetByID 方法
func (r *orderRepository) GetByID(id int) (*models.Order, error) {
	var order models.Order
	if err := r.db.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
