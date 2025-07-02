package models

import 	"time"

type Booking struct  {
    ID           uint      `json:"id" gorm:"primaryKey"`
    UserID       uint      `json:"user_id" gorm:"not null"`
    AttractionID uint      `json:"attraction_id" gorm:"not null"`
    Date         time.Time `json:"date" gorm:"not null"`
    Time         string    `json:"time" gorm:"size:10;not null"` // morning or afternoon
    Price        int       `json:"price" gorm:"not null"`
    CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
    
    User       User       `json:"user" gorm:"foreignKey:UserID"`
    Attraction Attraction `json:"attraction" gorm:"foreignKey:AttractionID"`
}

