// filepath: taipei-day-trip-go-go/internal/handlers/attraction.go
package handlers

import (
	"net/http"
	"taipei-day-trip-go-go/internal/interfaces"

	"github.com/gin-gonic/gin"
)

type AttractionHandler struct {
	Service interfaces.AttractionService
}

// NewAttractionHandler 現在需要傳入 *gorm.DB 實例
func NewAttractionHandler(service interfaces.AttractionService) *AttractionHandler {
	return &AttractionHandler{Service: service}
}

func (h *AttractionHandler) GetAttractions(c *gin.Context) {
	// 這裡可以根據需求取得 page/limit 參數
	page := 1
	limit := 10

	attractions, err := h.Service.ListAttractions(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法取得景點資料"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": attractions})
}
