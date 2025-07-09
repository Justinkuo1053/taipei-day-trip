package handlers

import (
	"net/http"
	"taipei-day-trip-go-go/internal/interfaces"
	"taipei-day-trip-go-go/internal/models"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	Service interfaces.OrderService
}

func NewOrderHandler(service interfaces.OrderService) *OrderHandler {
	return &OrderHandler{Service: service}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var input models.OrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.OrderCreateResponse{
			Error:   true,
			Message: "輸入資料格式錯誤",
		})
		return
	}

	orderNumber, err := h.Service.CreateOrder(input)
	if err != nil {
		c.JSON(http.StatusOK, models.OrderCreateResponse{
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
					ID:      order.Attraction.ID,
					Name:    order.Attraction.Name,
					Address: order.Attraction.Address,
					Image:   order.Attraction.Image,
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
