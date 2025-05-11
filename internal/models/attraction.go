package models

type Attraction struct {
    ID          uint     `json:"id" gorm:"primaryKey"`
    Name        string   `json:"name" gorm:"size:255;not null"`
    Description string   `json:"description" gorm:"type:text"`
    Address     string   `json:"address" gorm:"size:255"`
    Transport   string   `json:"transport" gorm:"type:text"`
    Category    string   `json:"category" gorm:"size:50"`
    Images      []string `json:"images" gorm:"type:json"`
    Latitude    float64  `json:"latitude"`
    Longitude   float64  `json:"longitude"`
}


