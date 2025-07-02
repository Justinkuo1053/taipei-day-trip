package models

import "time"

type Order struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	OrderNumber   string    `json:"order_number" gorm:"size:50;unique;not null"`
	UserID        uint      `json:"user_id" gorm:"not null"`
	BookingID     uint      `json:"booking_id" gorm:"not null"`
	Status        string    `json:"status" gorm:"size:20;not null"` // pending, paid, cancelled
	TotalAmount   int       `json:"total_amount" gorm:"not null"`
	ContactName   string    `json:"contact_name" gorm:"size:255"`
	ContactEmail  string    `json:"contact_email" gorm:"size:255"`
	ContactPhone  string    `json:"contact_phone" gorm:"size:20"`
	PaymentStatus string    `json:"payment_status" gorm:"size:20"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	User    User    `json:"user" gorm:"foreignKey:UserID"`
	Booking Booking `json:"booking" gorm:"foreignKey:BookingID"`
}

type OrderInput struct {
	// 定義欄位
}

type PaymentInput struct {
	// 定義欄位
}
