// filepath: taipei-day-trip-go-go/internal/models/attraction.go
package models

import "gorm.io/gorm"

type Attraction struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:255" json:"name"`
	Category    string `gorm:"size:255" json:"category"`
	Description string `gorm:"type:text" json:"description"`
	Address     string `gorm:"size:255" json:"address"`
	Transport   string `gorm:"size:text" json:"transport"`
	MRT         string `gorm:"size:255" json:"MRT"`
	Lat         float64
	Lng         float64
	Images      []string `gorm:"-" json:"images"` // API 回傳用，不存 DB
}

type Image struct {
	ID           uint   `gorm:"primaryKey"`
	URL          string `gorm:"size:512"`
	AttractionID uint
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Attraction{}, &Image{})
}
