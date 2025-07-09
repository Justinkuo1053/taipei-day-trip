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
