package handlers

import (
	"fmt"
	"net/http"
	"taipei-day-trip-go-go/internal/interfaces"
	"taipei-day-trip-go-go/internal/models"
	"taipei-day-trip-go-go/internal/services"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	Service        interfaces.OrderService
	BookingService *services.BookingService
}

func NewOrderHandler(service interfaces.OrderService, bookingService *services.BookingService) *OrderHandler {
	return &OrderHandler{Service: service, BookingService: bookingService}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	// 取得 userID
	userObj, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.OrderCreateResponse{
			Error:   true,
			Message: "未登入或無法取得使用者資訊",
		})
		return
	}
	user, ok := userObj.(*models.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, models.OrderCreateResponse{
			Error:   true,
			Message: "使用者資訊格式錯誤",
		})
		return
	}
	userID := user.ID

	// 解析訂單資料
	var input models.OrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.OrderCreateResponse{
			Error:   true,
			Message: "輸入資料格式錯誤",
		})
		return
	}
	if input.Prime == "" {
		c.JSON(http.StatusBadRequest, models.OrderCreateResponse{
			Error:   true,
			Message: "缺少 prime 或格式錯誤",
		})
		return
	}
	fmt.Println("收到 prime:", input.Prime)

	// 查詢該 user 的 booking，取得 bookingID
	booking, err := h.BookingService.GetBookingByUserID(userID)
	if err != nil || booking == nil {
		c.JSON(http.StatusBadRequest, models.OrderCreateResponse{
			Error:   true,
			Message: "找不到對應預定行程，請先預定行程",
		})
		return
	}

	// 呼叫 Service 建立訂單
	orderNumber, err := h.Service.CreateOrder(input, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.OrderCreateResponse{
			Error:   true,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.OrderCreateResponse{
		Data: &struct {
			Number  string               `json:"number"`
			Payment models.PaymentResult `json:"payment"`
		}{
			Number: orderNumber,
			Payment: models.PaymentResult{
				Status:  0,
				Message: "付款成功（mock）",
			},
		},
	})
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	orderNumber := c.Param("orderNumber")
	order, err := h.Service.GetOrder(orderNumber)
	if err != nil {
		c.JSON(http.StatusOK, models.OrderDetailResponse{
			Error:   true,
			Message: "查無此訂單",
		})
		return
	}
	// 組裝回傳格式
	resp := models.OrderDetailResponse{
		Data: &struct {
			Number  string             `json:"number"`
			Price   int                `json:"price"`
			Trip    models.TripInfo    `json:"trip"`
			Contact models.ContactInfo `json:"contact"`
			Status  int                `json:"status"`
		}{
			Number: order.OrderNumber,
			Price:  order.Price,
			Trip: models.TripInfo{
				Attraction: models.AttractionInfo{
					ID:      int(order.Attraction.ID),
					Name:    order.Attraction.Name,
					Address: order.Attraction.Address,
					Image:   "", // Attraction 沒有 Image 欄位，先給空字串或可自訂
				},
				Date: order.TripDate,
				Time: order.TripTime,
			},
			Contact: models.ContactInfo{
				Name:  order.ContactName,
				Email: order.ContactEmail,
				Phone: order.ContactPhone,
			},
			Status: order.Status,
		},
	}
	c.JSON(http.StatusOK, resp)
}
