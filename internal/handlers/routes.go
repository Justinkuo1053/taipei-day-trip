// filepath: taipei-day-trip-go-go/internal/handlers/routes.go
package handlers

import (
	"taipei-day-trip-go-go/internal/repositories"
	"taipei-day-trip-go-go/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, attractionHandler *AttractionHandler, bookingHandler *BookingHandler, userHandler *UserHandler, userService *services.UserService) {
	orderRepo := repositories.NewOrderRepositoryImpl()
	orderService := services.NewOrderServiceImpl(orderRepo)
	// 取得 bookingService
	// bookingHandler 已經有 BookingService 欄位
	orderHandler := NewOrderHandler(orderService, bookingHandler.BookingService)

	r.GET("/api/attractions", attractionHandler.GetAttractions)
	r.GET("/api/mrts", attractionHandler.GetMRTs)
	r.GET("/api/attraction/:id", attractionHandler.GetAttractionByID)
	r.GET("/api/attractions/search", attractionHandler.SearchAttractionsByKeyword) // 新增全文搜尋 API 路由

	r.Use(AuthMiddleware(userService))
	r.GET("/api/booking", bookingHandler.GetBooking)
	r.POST("/api/booking", bookingHandler.CreateBooking)
	r.DELETE("/api/booking/:id", bookingHandler.DeleteBooking)

	r.POST("/api/orders", orderHandler.CreateOrder)
	r.GET("/api/order/:orderNumber", orderHandler.GetOrder)

	r.POST("/api/user", userHandler.Register)
	r.PUT("/api/user/auth", userHandler.Login)
	r.GET("/api/user/auth", userHandler.GetAuthUser)
}
