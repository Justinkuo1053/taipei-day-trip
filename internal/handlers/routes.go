// filepath: taipei-day-trip-go-go/internal/handlers/routes.go
package handlers

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, attractionHandler *AttractionHandler) {
	r.GET("/api/attractions", attractionHandler.GetAttractions)
}
