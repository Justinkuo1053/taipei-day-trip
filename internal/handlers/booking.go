package handlers

import (
	"net/http"
	"strconv"
	"time"

	"taipei-day-trip-go-go/internal/models"
	"taipei-day-trip-go-go/internal/services"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	BookingService *services.BookingService
}

func NewBookingHandler(bookingService *services.BookingService) *BookingHandler {
	return &BookingHandler{BookingService: bookingService}
}

// 取得目前預定行程
func (h *BookingHandler) GetBooking(c *gin.Context) {
	// TODO: 取得 userID，這裡先寫死 1
	userID := uint(1)
	booking, err := h.BookingService.GetBookingByUserID(userID)
	if err != nil || booking == nil {
		c.JSON(http.StatusOK, gin.H{"data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"attraction": gin.H{
				"id":      booking.Attraction.ID,
				"name":    booking.Attraction.Name,
				"address": booking.Attraction.Address,
				"image":   "https://yourdomain.com/images/attraction/" + strconv.Itoa(int(booking.Attraction.ID)) + ".jpg", // TODO: 實際圖片
			},
			"date":  booking.Date.Format("2006-01-02"),
			"time":  booking.Time,
			"price": booking.Price,
		},
	})
}

// 建立新的預定行程
func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req struct {
		AttractionID uint   `json:"attractionId"`
		Date         string `json:"date"`
		Time         string `json:"time"`
		Price        int    `json:"price"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "請提供提供對應欄位的錯誤訊息."})
		return
	}
	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "日期格式錯誤"})
		return
	}
	userID := uint(1) // TODO: 取得 userID
	booking := &models.Booking{
		UserID:       userID,
		AttractionID: req.AttractionID,
		Date:         parsedDate,
		Time:         req.Time,
		Price:        req.Price,
	}
	err = h.BookingService.CreateBooking(booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "建立失敗"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

// 刪除單一預定行程
// 此 handler 會根據路由參數取得要刪除的訂單 id，並呼叫 service 層執行刪除
func (h *BookingHandler) DeleteBooking(c *gin.Context) {
	// 從路由參數取得訂單 id（如 /api/booking/3 會取得 "3"）
	idStr := c.Param("id")
	// 將 id 字串轉為整數型別，方便後續查詢
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// 若轉換失敗，回傳 400 Bad Request，表示訂單 id 格式錯誤
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "無效的訂單ID"})
		return
	}
	// TODO: 之後可加上 userID 權限驗證，確保只能刪除自己的訂單

	// 呼叫 service 層，根據 id 刪除單一訂單
	err = h.BookingService.DeleteBookingByID(uint(id))
	if err != nil {
		// 若刪除失敗（如資料庫錯誤或找不到該筆），回傳 500 Internal Server Error
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "刪除失敗"})
		return
	}
	// 刪除成功，回傳 ok
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
