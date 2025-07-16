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

// 依訂單編號查詢訂單，並預載 Attraction 與圖片
func (r *OrderRepositoryImpl) GetOrderByNumber(orderNumber string) (*models.Order, error) {
	var order models.Order
	// 預載 Attraction
	err := utils.Database.Preload("Attraction").Where("order_number = ?", orderNumber).First(&order).Error
	if err != nil {
		return nil, err
	}
	// 查詢 Attraction 的所有圖片，組成 Images 陣列
	var images []models.Image
	utils.Database.Where("attraction_id = ?", order.AttractionID).Find(&images)
	imageURLs := make([]string, 0, len(images))
	for _, img := range images {
		imageURLs = append(imageURLs, img.URL)
	}
	order.Attraction.Images = imageURLs
	return &order, nil
}
