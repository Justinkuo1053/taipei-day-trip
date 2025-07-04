// filepath: taipei-day-trip-go-go/internal/handlers/routes.go
package handlers

import (
	"taipei-day-trip-go-go/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, attractionHandler *AttractionHandler, bookingHandler *BookingHandler, userHandler *UserHandler, userService *services.UserService) {
	r.GET("/api/attractions", attractionHandler.GetAttractions)
	r.GET("/api/mrts", attractionHandler.GetMRTs)
	r.GET("/api/attractions/:id", attractionHandler.GetAttractionByID)

	r.GET("/api/booking", bookingHandler.GetBooking)
	r.POST("/api/booking", bookingHandler.CreateBooking)
	r.DELETE("/api/booking/:id", bookingHandler.DeleteBooking)
	r.Use(AuthMiddleware(userService))
	r.POST("/api/user", userHandler.Register)
	r.PUT("/api/user/auth", userHandler.Login)
	r.GET("/api/user/auth", userHandler.GetAuthUser)
}
