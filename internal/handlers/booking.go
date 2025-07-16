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
	// 取得 userID，從 JWT middleware 取得
	userObj, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"data": nil, "error": true, "message": "未登入或無法取得使用者資訊"})
		return
	}
	user, ok := userObj.(*models.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"data": nil, "error": true, "message": "使用者資訊格式錯誤"})
		return
	}
	userID := user.ID

	booking, err := h.BookingService.GetBookingByUserID(userID)
	if err != nil || booking == nil {
		c.JSON(http.StatusOK, gin.H{"data": nil})
		return
	}

	// 回傳 booking 與景點資料（從 booking.Attraction 取得）
	attr := booking.Attraction
	imageUrl := ""
	if len(attr.Images) > 1 {
		imageUrl = attr.Images[1]
	} else if len(attr.Images) > 0 {
		imageUrl = attr.Images[0]
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"attraction": gin.H{
				"id":      attr.ID,
				"name":    attr.Name,
				"address": attr.Address,
				"image":   imageUrl,
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
	// 取得 userID，從 JWT middleware 取得
	userObj, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "未登入或無法取得使用者資訊"})
		return
	}
	user, ok := userObj.(*models.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "使用者資訊格式錯誤"})
		return
	}
	userID := user.ID
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
