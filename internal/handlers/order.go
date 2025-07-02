package handlers

// import (
// 	"net/http"
// 	"taipei-day-trip-go-go/internal/interfaces"
// 	"taipei-day-trip-go-go/internal/models"

// 	"github.com/gin-gonic/gin"
// )

// type OrderHandler struct {
// 	Service interfaces.OrderService
// }

// func NewOrderHandler(service interfaces.OrderService) *OrderHandler {
// 	return &OrderHandler{Service: service}
// }

// func (h *OrderHandler) CreateOrder(c *gin.Context) {
// 	var input models.OrderInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	orderNumber, err := h.Service.CreateOrder(input)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"order_number": orderNumber})
// }

