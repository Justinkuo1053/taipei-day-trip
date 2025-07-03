// filepath: taipei-day-trip-go-go/internal/models/attraction.go
package models

import "gorm.io/gorm"

type Attraction struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255"`
	Category    string `gorm:"size:255"`
	Description string `gorm:"type:text"`
	Address     string `gorm:"size:255"`
	Transport   string `gorm:"size:text"`
	MRT         string `gorm:"size:255"`
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
